package bot

import (
	"errors"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/kaverhovsky/gosniias-time-manager-bot/pkg/common"
	"go.uber.org/zap"
)

type Bot struct {
	api    *tgbotapi.BotAPI
	config *common.Config
	logger *zap.Logger
	done   chan struct{}
}

func New(logger *zap.Logger, config *common.Config) *Bot {
	return &Bot{
		logger: logger,
		config: config,
		done:   make(chan struct{}),
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

func (b *Bot) Listen() {
	updates := b.api.ListenForWebhook("/" + b.config.BotToken)
	b.logger.Debug("started processing loop")
	for {
		select {
		case update := <-updates:
			b.logger.With(zap.String("message", update.Message.Text)).Info("received a message")
		case <-b.done:
			return
		}
	}
}
