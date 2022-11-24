package usecase

import (
	"github.com/kaverhovsky/gosniias-time-manager-bot/internal/repository"
	"github.com/kaverhovsky/gosniias-time-manager-bot/pkg/common"
	"go.uber.org/zap"
)

type Usecase struct {
	logger *zap.Logger
	config *common.Config
	repo   repository.Repository
}

func New(logger *zap.Logger, config *common.Config, repo repository.Repository) *Usecase {
	return &Usecase{
		logger: logger,
		config: config,
		repo:   repo,
	}
}

func Get() {

}
