package main

import (
	"flag"
	"github.com/kaverhovsky/gosniias-time-manager-bot/internal/bot"
	"github.com/kaverhovsky/gosniias-time-manager-bot/pkg/common"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	configPath := flag.String("c", "", "config file path")
	flag.Parse()

	startUpLogger := common.NewLogger("development", "info")
	config, err := common.ReadConfig(*configPath, startUpLogger)
	if err != nil {
		startUpLogger.With(zap.String("reason", err.Error())).Error("failed config read")
		return
	}
	logger := common.NewLogger(config.Mode(), config.LogLevel)

	bot := bot.New(logger, config)
	if err := bot.Init(config.BotToken); err != nil {
		logger.With(zap.String("reason", err.Error())).Error("failed initialising bot")
		return
	}

	go bot.RunForUpdates()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM, syscall.SIGINT)
	<-sigs
	bot.Shutdown()
	logger.Info("service stopped")
}
