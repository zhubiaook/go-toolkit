package main

import (
	"log/slog"
	"os"
)

type LoggerOptions struct {
	// 是否开启日志源记录，如果开启会在日志中显示调用日志所在的文件和行号
	AddSource bool
	// 指定日志级别，可选值：debug, info, warn, error
	Level string
	// 指定日志显示格式，可选值：text, json
	Format string
	// 指定日志输出位置
	OutputPath string
}

func InitLogger(o LoggerOptions) {
	var level slog.Level
	switch o.Level {
	case "debug":
		level = slog.LevelDebug
	case "info":
		level = slog.LevelInfo
	case "warn":
		level = slog.LevelWarn
	case "error":
		level = slog.LevelError
	default:
		level = slog.LevelInfo
	}

	opts := &slog.HandlerOptions{
		AddSource: o.AddSource,
		Level:     level,
	}

	var handler slog.Handler
	switch o.Format {
	case "json":
		handler = slog.NewJSONHandler(os.Stdout, opts)
	case "text":
		handler = slog.NewTextHandler(os.Stdout, opts)
	default:
		handler = slog.NewTextHandler(os.Stdout, opts)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)
}
