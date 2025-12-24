package main

import (
    "strconv"
    "time"
    "context"
    "fmt"
    "log"
    "os"
    "encoding/json"
    "github.com/go-redis/redis/v8"
    "hermawan-monitora/hmonglobal/lang"
    "hermawan-monitora/module/hmonenv"
    "hermawan-monitora/module/hmonredis"
    "net"
)

type ServerPort struct {
  Ip string
  Port int
  ServerName string
  ServiceName string
  ServerGroup string
}

var ctx = context.Background()
var portScanStatus = "INIT"

func getAllServerNPorts(oldServerPorts []ServerPort) []ServerPort {
    var err error
    var result string
    for _, val := range oldServerPorts {
        err = hmonredis.Del(hmonredis.GetMPPubSubServerNPort(val.Ip, val.Port))
        if err != nil {
            panic(fmt.Sprintf("ERROR getAllServerNPorts Del %s\n", err.Error()))
        }
    }
    var out []ServerPort
    result, err = hmonredis.Get(hmonredis.GetMPAllServerPort())
    if err != nil {
        panic(fmt.Sprintf("ERROR getAllServerNPorts 1 %s\n", err.Error()))
    }
    var allServerNPorts []map[string]interface{}
    if err = json.Unmarshal([]byte(result), &allServerNPorts); err != nil {
        log.Println(fmt.Sprintf("Result = %v\n", result))
        panic(fmt.Sprintf("ERROR getAllServerNPorts 2 %s\n", lang.JsonCnvErr(err.Error())))
    }
    for _, val := range allServerNPorts {
        var temp ServerPort
        temp.Ip = val["ip"].(string)
        temp.Port = int(val["port"].(float64))
        temp.ServerName = val["serverName"].(string)
        temp.ServiceName = val["serviceName"].(string)
        temp.ServerGroup = val["serverGroup"].(string)
        out = append(out, temp)
    }
    strBytes := []byte("SCAN")
    err = hmonredis.Publish(hmonredis.MPPortScanStatus, strBytes)
    if err != nil {
        panic(fmt.Sprintf("ERROR getAllServerNPorts 3 %s\n", err.Error()))
    }
    return out
}

func scanProcess(resetProcess chan bool) {
    var microSeconds int64
    lastStatus := make(map[string]map[int]string)
    serverPorts := getAllServerNPorts([]ServerPort{})
    for {
        select {
        case <-resetProcess:
            serverPorts = getAllServerNPorts(serverPorts)
        default:
            for _, val := range serverPorts {
                microSeconds = time.Now().UnixMicro()
                ip := val.Ip + ":" + strconv.Itoa(val.Port)
                _, err := net.DialTimeout("tcp", ip, time.Duration(300)*time.Millisecond)
                status := "DOWN"
                if err == nil {
                    status = "UP"
                    log.Println(fmt.Printf("%s is Available!", val.ServerName))
                }
                tempLastStatus, ok := lastStatus[val.Ip][val.Port]
                if ok && (tempLastStatus == status) {

                } else {
                    _, ok2 := lastStatus[val.Ip]
                    if !ok2 {
                        lastStatus[val.Ip] = make(map[int]string)
                    }
                    lastStatus[val.Ip][val.Port] = status
                    packet := map[string]any {
                      "status": status,
                      "time": microSeconds}
                    jsonInBytes, err := json.Marshal(&packet)
                    if err != nil {
                        panic(fmt.Sprintf("ERROR 1 %s\n", err.Error()))
                    }
                    err = hmonredis.Publish(hmonredis.GetMPPubSubServerNPort(val.Ip, val.Port), jsonInBytes)
                    if err != nil {
                        panic(fmt.Sprintf("ERROR 2 %s\n", err.Error()))
                    }
                    err = hmonredis.SetStr(hmonredis.GetMPLastServerPortStatus(val.Ip, val.Port), status)
                    if err != nil {
                        panic(fmt.Sprintf("ERROR 3 %s\n", err.Error()))
                    }
                }
            }
        }
        time.Sleep(hmonenv.GetInterval())
    }
}

func main() {
    // Check Arguments
    args := os.Args
    if (len(args) == 2) && ((args[1] == "-h") || (args[1] == "--help")) {
        fmt.Printf(lang.Help())
        return
    }

    // Check bgserver
    bgServerStatus := ""
    var err error
    for bgServerStatus != "READY" {
        bgServerStatus, err = hmonredis.Get(hmonredis.BgServerStatus)
        if (err != nil) && (err.Error() != "EOF") {
            panic(fmt.Sprintf("ERROR Read Background-Server Status in Redis\nERROR = \"%s\"\n", err.Error()))
        }
    }

    // Process
    resetProcess := make(chan bool)
    go scanProcess(resetProcess)
    resetProcess <- true
    subscriber := hmonredis.Subscribe(ctx, hmonredis.MPPortScanStatus)
    var msg *redis.Message
    for {
        msg, err = hmonredis.SubscriberReceiveMessage(subscriber)
        if err != nil {
            panic(fmt.Sprintf("ERROR Redis subscribe\nERROR = \"%s\"\n", msg, err))
        }
	portScanStatus := string(msg.Payload)
        log.Println(fmt.Sprintf("REDIS status = %s", portScanStatus))
        if portScanStatus == "INIT" {
            resetProcess <- true
        }
    }
}
