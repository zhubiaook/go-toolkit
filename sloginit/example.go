package sloginit

import (
	"log"
	"log/slog"

	"github.com/spf13/viper"
)

func init() {
	cfgFile := "./config.yaml"
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("ReadInConfig failed: %v\n", err)
	}
}

func example() {
	loggerOptions := LoggerOptions{
		AddSource:  viper.GetBool("log.add-source"),
		Level:      viper.GetString("log.level"),
		Format:     viper.GetString("log.format"),
		OutputPath: viper.GetString("log.output-path"),
	}

	InitLogger(loggerOptions)

	slog.Debug("debug")
	slog.Info("info")
	slog.Warn("warn")
	slog.Error("error")
}
