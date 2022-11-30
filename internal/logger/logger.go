package logger

import (
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Writer io.Writer = os.Stderr

// InitWriter 初始化日志输出 writer
func InitWriter() {
	if logPath := viper.GetString("log.path"); logPath != "" { // 设置滚动写入文件
		if err := os.MkdirAll(logPath, 0777); err != nil {
			panic(err)
		}

		Writer = &lumberjack.Logger{
			Filename:   filepath.Join(logPath, "server.log"),
			MaxSize:    500, // megabytes
			MaxBackups: 10,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		}
	}
}

// Init 初始化日志组件
func Init() {
	InitWriter()

	opts := slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.Level(viper.GetInt("log.level")),
		ReplaceAttr: nil,
	}
	slog.SetDefault(slog.New(opts.NewJSONHandler(Writer)))
}
