// examples/main.go
package main

import (
    "github.com/your-username/common-lib/pkg/encrypt"
    "github.com/your-username/common-lib/pkg/logger"
    "log"
)

func main() {
    // 使用密码加密
    pwd := encrypt.NewBCryptPassword(0)
    hashed, err := pwd.Hash("123456")
    if err != nil {
        log.Fatal(err)
    }

    // 验证密码
    err = pwd.Compare(hashed, "123456")
    if err != nil {
        log.Fatal(err)
    }

    // 初始化日志
    logger, err := logger.NewLogger(logger.Config{
        Level:    "info",
        Filename: "app.log",
    })
    if err != nil {
        log.Fatal(err)
    }
    defer logger.Sync()
}