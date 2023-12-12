package log

import (
	"context"
	"fmt"
	"log"
	"os"
)

func InfoContextf(ctx context.Context, format string, args ...interface{}) {
	if DefaultLogger == nil {
		return
	}
	DefaultLogger.InfoContextf(ctx, format, args...)
}

// 默认日志输出，可以自己注入
var DefaultLogger Logger = &logger{l: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)}

type Logger interface {
	InfoContextf(ctx context.Context, format string, args ...interface{})
}
type logger struct {
	l *log.Logger
}

func (l *logger) InfoContextf(ctx context.Context, format string, args ...interface{}) {
	format = fmt.Sprintf("[INFO] %s\n", format)
	l.l.Printf(format, args...)
}
