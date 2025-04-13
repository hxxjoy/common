package logger

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"
)

type Config struct {
	Level      string
	Filename   string
	MaxSize    int  // 单个文件最大尺寸，单位MB
	MaxBackups int  // 最大保留文件数
	MaxAge     int  // 最大保留天数
	Compress   bool // 是否压缩
}

type Logger struct {
	config     *Config
	file       *os.File
	size       int64
	createTime time.Time
}

func NewLogger(cfg Config) (*Logger, error) {
	logger := &Logger{
		config: &cfg,
	}

	if err := logger.rotate(); err != nil {
		return nil, err
	}

	return logger, nil
}

func (l *Logger) Write(p []byte) (n int, err error) {
	writeLen := int64(len(p))

	if l.size+writeLen > int64(l.config.MaxSize*1024*1024) {
		if err := l.rotate(); err != nil {
			return 0, err
		}
	}

	n, err = l.file.Write(p)
	l.size += int64(n)
	return n, err
}

func (l *Logger) rotate() error {
	if l.file != nil {
		l.file.Close()
	}

	// 生成新的文件名
	filename := l.config.Filename
	if l.shouldRotate() {
		now := time.Now()
		timestamp := now.Format("2006-01-02_15-04-05")
		filename = fmt.Sprintf("%s.%s", l.config.Filename, timestamp)
	}

	// 创建新文件
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	l.file = file
	l.size = 0
	l.createTime = time.Now()

	// 清理旧文件
	if err := l.cleanup(); err != nil {
		return err
	}

	return nil
}

func (l *Logger) shouldRotate() bool {
	if l.file == nil {
		return false
	}

	if l.size >= int64(l.config.MaxSize*1024*1024) {
		return true
	}

	return false
}

func (l *Logger) cleanup() error {
	if l.config.MaxBackups == 0 && l.config.MaxAge == 0 {
		return nil
	}

	pattern := l.config.Filename + ".*"
	matches, err := filepath.Glob(pattern)
	if err != nil {
		return err
	}

	// 按时间排序
	sort.Slice(matches, func(i, j int) bool {
		return matches[i] > matches[j]
	})

	// 处理文件数量限制
	for i := l.config.MaxBackups; i < len(matches); i++ {
		if err := os.Remove(matches[i]); err != nil {
			return err
		}
	}

	// 处理文件年龄限制
	if l.config.MaxAge > 0 {
		cutoff := time.Now().Add(-time.Duration(l.config.MaxAge) * 24 * time.Hour)
		for _, match := range matches {
			t, err := time.Parse("2006-01-02_15-04-05", strings.TrimPrefix(match, l.config.Filename+"."))
			if err != nil {
				continue
			}
			if t.Before(cutoff) {
				if err := os.Remove(match); err != nil {
					return err
				}
			}
		}
	}

	// 处理压缩
	if l.config.Compress {
		for _, match := range matches {
			if !strings.HasSuffix(match, ".gz") {
				if err := l.compressFile(match); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (l *Logger) compressFile(filename string) error {
	in, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(filename + ".gz")
	if err != nil {
		return err
	}
	defer out.Close()

	gz := gzip.NewWriter(out)
	defer gz.Close()

	if _, err := io.Copy(gz, in); err != nil {
		return err
	}

	// 压缩完成后删除原文件
	if err := os.Remove(filename); err != nil {
		return err
	}

	return nil
}

// Close 关闭日志文件
func (l *Logger) Close() error {
	if l.file != nil {
		return l.file.Close()
	}
	return nil
}
