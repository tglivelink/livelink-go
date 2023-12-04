package log

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tglivelink/livelink-go/pkg/util"
)

func Infof(ctx context.Context, format string, args ...interface{}) {
	if DefaultLogger == nil {
		return
	}
	DefaultLogger.Infof(ctx, format, args...)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	if DefaultLogger == nil {
		return
	}
	DefaultLogger.Errorf(ctx, format, args...)
}

// 默认日志输出，可以自己注入
var DefaultLogger Logger = &logger{l: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds)}

type Logger interface {
	Infof(ctx context.Context, format string, args ...interface{})
	Errorf(ctx context.Context, format string, args ...interface{})
}
type logger struct {
	l *log.Logger
}

func (l *logger) Infof(ctx context.Context, format string, args ...interface{}) {
	format = fmt.Sprintf("[INFO] %s, RID:%s\n", format, util.TraceID(ctx))
	l.l.Printf(format, args...)
}

func (l *logger) Errorf(ctx context.Context, format string, args ...interface{}) {
	format = fmt.Sprintf("[ERROR] %s, RID:%s\n", format, util.TraceID(ctx))
	l.l.Printf(format, args...)
}
