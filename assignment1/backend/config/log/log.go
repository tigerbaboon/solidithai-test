package log

import (
	"context"
	"fmt"
	"log/slog"
	"runtime"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/exp/zapslog"
)

func Init() {
	var zLog *zap.Logger
	var err error
	zLog, err = zap.NewDevelopment()
	if err != nil {
		panic(err)
	}

	slog.SetDefault(slog.New(zapslog.NewHandler(zLog.Core(), &zapslog.HandlerOptions{
		AddSource:  true,
		LoggerName: viper.GetString("APP_NAME"),
	})))
}

type Logger struct {
	*slog.Logger
}

func Default() *Logger {
	return &Logger{slog.Default()}
}

func (l *Logger) log(ctx context.Context, level slog.Level, msg string, args ...any) {
	if !l.Enabled(ctx, level) {
		return
	}
	var pc uintptr
	var pcs [1]uintptr
	runtime.Callers(3, pcs[:])
	pc = pcs[0]
	r := slog.NewRecord(time.Now(), level, msg, pc)
	r.Add(args...)
	if ctx == nil {
		ctx = context.Background()
	}
	_ = l.Handler().Handle(ctx, r)
}

func Debug(format string, args ...any) {
	Default().log(context.Background(), slog.LevelDebug, fmt.Sprintf(format, args...))
}
func Info(format string, args ...any) {
	Default().log(context.Background(), slog.LevelInfo, fmt.Sprintf(format, args...))
}

func Warn(format string, args ...any) {
	Default().log(context.Background(), slog.LevelWarn, fmt.Sprintf(format, args...))
}

func Error(format string, args ...any) {
	Default().log(context.Background(), slog.LevelError, fmt.Sprintf(format, args...))
}
