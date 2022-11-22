package bot

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"go.uber.org/zap"
)

type Bot struct {
	api    *tgbotapi.BotAPI
	logger *zap.Logger
}

func New(logger *zap.Logger) *Bot {
	return &Bot{
		logger: logger,
	}
}

func (b *Bot) Init(token string) error {
	if token == "" {
		return errors.New("telegram bot token string is empty")
	}
	api, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		b.logger.Error("failed creating new bot api")
		return err
	}
	b.api = api
	return nil
}
