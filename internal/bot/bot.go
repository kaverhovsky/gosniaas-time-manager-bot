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
		logger: logger.Named("bot"),
		config: config,
		done:   make(chan struct{}),
	}
}

func (b *Bot) Init(token string) error {
	if token == "" {
		return errors.New("telegram bot token string is empty")
	}
	api, err := tgbotapi.NewBotAPI(token)
	if b.config.Debug {
		api.Debug = true
	}

	if err != nil {
		b.logger.Error("failed creating new bot api")
		return err
	}
	b.api = api
	return nil
}

func (b *Bot) RunForUpdates() {
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30
	updates := b.api.GetUpdatesChan(updateConfig)

	for {
		select {
		case update := <-updates:
			if update.Message == nil {
				continue
			}
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
			if _, err := b.api.Send(msg); err != nil {
				b.logger.With(zap.String("reason", err.Error())).Error("failed message send")
			}
		case <-b.done:
			b.logger.Info("canceling bot run for updates")
			return
		}
	}

}

func (b *Bot) ListenForWebhook() {
	updates := b.api.ListenForWebhook("/" + b.config.BotToken)
	b.logger.Debug("started processing loop")
	for {
		select {
		case update := <-updates:
			b.logger.With(zap.String("message", update.Message.Text)).Info("received a message")
		case <-b.done:
			b.logger.Info("canceling bot listen")
			return
		}
	}
}

func (b *Bot) Shutdown() {
	close(b.done)
}
