package lang

import (
    "fmt"
)

func Help() string {
    return `usage: hermawan-monitora-worker-mp [-h] [--env ENV]

Hermawan-Monitora is a monitoring application, and the main feature is for monitoring services' health.
This is the background service.
For more description and help please check:
https://github.com/ming-hermawan/hermawan-monitora-manual/blob/master/hemawan-monitora-manual.pdf

optional arguments:
  -h, --help  show this help message and exit
  -env        env file path
`
}

func DbConnErr(errMsg string) string {
    return fmt.Sprintf("Error database connection! %s", errMsg)
}

func ReadDbErr(lbl string, errMsg string) string {
    return fmt.Sprintf("Error read %s from database! %s", lbl, errMsg)
}

func ReadRedisErr(lbl string, errMsg string) string {
    return fmt.Sprintf("Error read %s from Redis! %s", lbl, errMsg)
}

func JsonCnvErr(errMsg string) string {
    return fmt.Sprintf("Convert to JSON error! %s", errMsg)
}
