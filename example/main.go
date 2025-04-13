// examples/main.go
package main

import (
	"log"

	"github.com/hxxjoy/common/pkg/encrypt"
	"github.com/hxxjoy/common/pkg/logger"
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
		Level:      "info",
		Filename:   "./logs/app.log",
		MaxSize:    100,  // 100MB
		MaxBackups: 3,    // 保留3个备份
		MaxAge:     7,    // 保留7天
		Compress:   true, // 压缩旧文件
	})
	if err != nil {
		panic(err)
	}
	defer logger.Close()

	// 使用logger写入日志
	logger.Write([]byte("这是一条日志\n"))
}
