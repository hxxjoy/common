// pkg/logger/zap.go
package logger

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
)

// Config 日志配置
type Config struct {
    Level      string
    Filename   string
    MaxSize    int
    MaxBackups int
    MaxAge     int
    Compress   bool
}

// NewLogger 创建日志实例
func NewLogger(cfg Config) (*zap.Logger, error) {
    // 日志配置实现...
}