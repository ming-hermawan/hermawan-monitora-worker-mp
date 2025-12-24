package hmonenv

import (
    "fmt"
    "os"
    "time"
    "strconv"
    "strings"
    "crypto/tls"
    "github.com/joho/godotenv"
)

var envFilePath = "/etc/hermawan-monitora/worker-mp.env"
var interval time.Duration
var logDirPath string
var redisHost string
var redisPwd string
var redisPort int
var redisDb int
var redisMaxRetries int
var redisMinRetryBackoff time.Duration
var redisMaxRetryBackoff time.Duration
var redisDialTimeout time.Duration
var redisReadTimeout time.Duration
var redisWriteTimeout time.Duration
var redisPoolSize int
var redisMinIdleConns int
var redisMaxConnAge time.Duration
var redisPoolTimeout time.Duration
var redisIdleTimeout time.Duration
var redisIdleCheckFrequency time.Duration
var redisTlsConfig *tls.Config

func GetLogDirPath() string {
    return logDirPath
}

func GetRedisHost() string {
    return redisHost
}

func GetRedisPwd() string {
    return redisPwd
}

func GetRedisPort() int {
    return redisPort
}

func GetRedisDb() int {
    return redisDb
}

func GetRedisMaxRetries() int {
    return redisMaxRetries
}

func GetRedisMinRetryBackoff() time.Duration {
    return redisMinRetryBackoff
}

func GetRedisMaxRetryBackoff() time.Duration {
    return redisMaxRetryBackoff
}

func GetRedisDialTimeout() time.Duration {
    return redisDialTimeout
}

func GetRedisReadTimeout() time.Duration {
    return redisReadTimeout
}

func GetRedisWriteTimeout() time.Duration {
    return redisWriteTimeout
}

func GetRedisPoolSize() int {
    return redisPoolSize
}
func GetRedisMinIdleConns() int {
    return redisMinIdleConns
}
func GetRedisMaxConnAge() time.Duration {
    return redisMaxConnAge
}
func GetRedisPoolTimeout() time.Duration {
    return redisPoolTimeout
}
func GetRedisIdleTimeout() time.Duration {
    return redisIdleTimeout
}
func GetRedisIdleCheckFrequency() time.Duration {
    return redisIdleCheckFrequency
}
func GetRedisTlsConfig() *tls.Config {
    return redisTlsConfig
}

func GetInterval() time.Duration {
    return interval
}

func init() {

    // INITIALIZATION
    var found bool
    var err error
    var tempStr string
    var tempInt64 int64
    args := os.Args[1:]
    if len(args) > 1 {
        envFilePath = args[1]
    }
    err = godotenv.Load(envFilePath)
    if err != nil {
        panic(fmt.Sprintf("Error read '%s'!", envFilePath))
    }

    // INTERVAL
    tempStr, found = os.LookupEnv("INTERVAL")
    if found {
        tempInt64, err = strconv.ParseInt(tempStr, 10,  64)
        if err != nil {
            panic("INTERVAL error in env!")
        }
        interval = CnvToTimeDurationSeconds(tempInt64)
    } else {
        interval = CnvToTimeDurationSeconds(30)
    }

    // REDIS
    redisHost = os.Getenv("REDIS_HOST")
    redisPwd = os.Getenv("REDIS_PASSWORD")
    redisPort, err = strconv.Atoi(os.Getenv("REDIS_PORT"))
    if err != nil {
        panic("REDIS_PORT not integer in env!")
    }
    redisDb, err = strconv.Atoi(os.Getenv("REDIS_DB"))
    if err != nil {
        panic("REDIS_DB not integer in env!")
    }
    tempStr, found = os.LookupEnv("REDIS_MAX_RETRIES")
    if found {
        redisMaxRetries, err = strconv.Atoi(tempStr)
        if err != nil {
            panic("REDIS_MAX_RETRIES not integer in env!")
        }
    } else {
        redisMaxRetries = 0
    }
    tempStr, found = os.LookupEnv("REDIS_MIN_RETRY_BACKOFF")
    if found {
        tempInt64, err = strconv.ParseInt(tempStr, 10,  64)
        if err != nil {
            panic("REDIS_MIN_RETRY_BACKOFF not integer in env!")
        }
        redisMinRetryBackoff = CnvToTimeDurationMilliSeconds(tempInt64)
    } else {
        redisMinRetryBackoff = CnvToTimeDurationMilliSeconds(8)
    }
    tempStr, found = os.LookupEnv("REDIS_MAX_RETRY_BACKOFF")
    if found {
        tempInt64, err = strconv.ParseInt(tempStr, 10,  64)
        if err != nil {
            panic("REDIS_MAX_RETRY_BACKOFF not integer in env!")
        }
        redisMaxRetryBackoff = CnvToTimeDurationMilliSeconds(tempInt64)
    } else {
        redisMaxRetryBackoff = CnvToTimeDurationMilliSeconds(512)
    }
    tempStr, found = os.LookupEnv("REDIS_DIAL_TIMEOUT")
    if found {
        tempInt64, err = strconv.ParseInt(tempStr, 10,  64)
        if err != nil {
            panic("REDIS_DIAL_TIMEOUT error in env!")
        }
        redisDialTimeout = CnvToTimeDurationSeconds(tempInt64)
    } else {
        redisDialTimeout = CnvToTimeDurationSeconds(5)
    }
    tempStr, found = os.LookupEnv("REDIS_READ_TIMEOUT")
    if found {
        tempInt64, err = strconv.ParseInt(tempStr, 10,  64)
        if err != nil {
            panic("REDIS_READ_TIMEOUT error in env!")
        }
        redisReadTimeout = CnvToTimeDurationSeconds(tempInt64)
    } else {
        redisReadTimeout = CnvToTimeDurationSeconds(3)
    }
    tempStr, found = os.LookupEnv("REDIS_WRITE_TIMEOUT")
    if found {
        tempInt64, err = strconv.ParseInt(tempStr, 10,  64)
        if err != nil {
            panic("REDIS_WRITE_TIMEOUT error in env!")
        }
        redisWriteTimeout = CnvToTimeDurationSeconds(tempInt64)
    } else {
        redisWriteTimeout = redisReadTimeout
    }
    tempStr, found = os.LookupEnv("REDIS_POOL_SIZE")
    if found {
        redisPoolSize, err = strconv.Atoi(tempStr)
        if err != nil {
            panic("REDIS_POOL_SIZE not integer in env!")
        }
    } else {
        redisPoolSize = 10
    }
    tempStr, found = os.LookupEnv("REDIS_MIN_IDLE_CONNS")
    if found {
        redisMinIdleConns, err = strconv.Atoi(tempStr)
        if err != nil {
            panic("REDIS_MIN_IDLE_CONNS not integer in env!")
        }
    } else {
        redisMinIdleConns = 0
    }
    tempStr, found = os.LookupEnv("REDIS_MAX_CONN_AGE")
    if found {
        tempInt64, err = strconv.ParseInt(tempStr, 10,  64)
        if err != nil {
            panic("REDIS_MAX_CONN_AGE error in env!")
        }
        redisMaxConnAge = CnvToTimeDurationSeconds(tempInt64)
    } else {
        redisMaxConnAge = 0
    }
    tempStr, found = os.LookupEnv("REDIS_POOL_TIMEOUT")
    if found {
        tempInt64, err = strconv.ParseInt(tempStr, 10,  64)
        if err != nil {
            panic("REDIS_POOL_TIMEOUT error in env!")
        }
        redisPoolTimeout = CnvToTimeDurationSeconds(tempInt64)
    } else {
        redisPoolTimeout = redisReadTimeout + (time.Duration(1) * time.Second)
    }
    tempStr, found = os.LookupEnv("REDIS_IDLE_TIMEOUT")
    if found {
        tempInt64, err = strconv.ParseInt(tempStr, 10,  64)
        if err != nil {
            panic("REDIS_IDLE_TIMEOUT error in env!")
        }
        redisIdleTimeout = CnvToTimeDurationSeconds(tempInt64)
    } else {
        redisPoolTimeout = CnvToTimeDurationSeconds(300)
    }
    tempStr, found = os.LookupEnv("REDIS_IDLE_CHECK_FREQUENCY")
    if found {
        tempInt64, err = strconv.ParseInt(tempStr, 10,  64)
        if err != nil {
            panic("REDIS_IDLE_CHECK_FREQUENCY error in env!")
        }
        redisIdleCheckFrequency = CnvToTimeDurationSeconds(tempInt64)
    } else {
        redisIdleCheckFrequency = CnvToTimeDurationSeconds(60)
    }
    // TLS
    redisClientCrtPath, redisClientCrtPathFound := os.LookupEnv("REDIS_CLIENT_CRT_FILEPATH")
    redisClientKeyPath, redisClientKeyPathFound := os.LookupEnv("REDIS_CLIENT_KEY_FILEPATH")
    redisClientCaPath, redisClientCaPathFound := os.LookupEnv("REDIS_CLIENT_CA_FILEPATH")
    if strings.TrimSpace(redisClientCrtPath) == "" {
        redisClientCrtPath = ""
        redisClientCrtPathFound = false
    }
    if strings.TrimSpace(redisClientKeyPath) == "" {
        redisClientKeyPath = ""
        redisClientKeyPathFound = false
    }
    if strings.TrimSpace(redisClientCaPath) == "" {
        redisClientCaPath = ""
        redisClientCaPathFound = false
    }
    if redisClientCrtPathFound || redisClientKeyPathFound || redisClientCaPathFound {
        if !redisClientCrtPathFound && redisClientKeyPathFound {
            panic("REDIS_CLIENT_CRT_FILEPATH is not set in env!")
        }
        if !redisClientKeyPathFound && redisClientCrtPathFound {
            panic("REDIS_CLIENT_KEY_FILEPATH is not set in env!")
        }
        redisTlsConfig, tempStr = getTlsConfig(
          redisClientCrtPath,
	  redisClientKeyPath,
	  redisClientCaPath)
        if tempStr != "" {
            panic(tempStr)
        }
    }

    // PATH
    // Log
    logDirPath = os.Getenv("LOG_DIRPATH")
    _, err = os.Stat(logDirPath)
    if err != nil {
        panic(fmt.Sprintf("%s folder is not exists!", logDirPath))
    }
}
