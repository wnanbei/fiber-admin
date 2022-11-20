package logger

import (
	"io"
	"os"
	"path/filepath"

	"github.com/spf13/viper"
	"golang.org/x/exp/slog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Init 初始化日志组件
func Init() {
	opts := slog.HandlerOptions{
		AddSource:   true,
		Level:       slog.Level(viper.GetInt("log.level")),
		ReplaceAttr: nil,
	}

	var out io.Writer = os.Stderr

	if logPath := viper.GetString("log.path"); logPath != "" { // 设置滚动写入文件
		if err := os.MkdirAll(logPath, 0777); err != nil {
			panic(err)
		}

		out = &lumberjack.Logger{
			Filename:   filepath.Join(logPath, "server.log"),
			MaxSize:    500, // megabytes
			MaxBackups: 10,
			MaxAge:     28,   //days
			Compress:   true, // disabled by default
		}
	}

	slog.SetDefault(slog.New(opts.NewJSONHandler(out)))
}
