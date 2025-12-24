package hmonredis

import (
    "fmt"
    "github.com/google/uuid"
)

const (
    LogoutAfter1Hour = "logout-after-1-hour"
    MPPortScanStatus = "mp-ports-scan-status"
    BgServerStatus = "bg-server-status"
)

func GetFailedLoginKey(username string) string {
    return fmt.Sprintf("failed-login-%s", username)
}

func GetMPLastServerPortStatus(ip string, port int) string {
    return fmt.Sprintf("mp-last:%s:%d", ip, port)
}

func GetMonPortUploadCSVKey() string {
    return fmt.Sprintf("mp-csv-%s", uuid.New())
}

func GetMPPubSubServerNPort(ip string, port int) string {
    return fmt.Sprintf("mp-pubsub:%s:%d", ip, port)
}

func GetMPServerMails(ip string) string {
    return fmt.Sprintf("mp-mail:%s", ip)
}

func GetMPServiceName(ip string, port int) string {
    return fmt.Sprintf("mp-service-name:%s:%d", ip, port)
}

func GetMPAllServerPort() string {
    return fmt.Sprintf("mp-all-server-port")
}

func GetUsrMenu(username string, menu string) string {
    return fmt.Sprintf("menu-%s-%s", username, menu)
}
