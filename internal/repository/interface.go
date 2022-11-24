package repository

import (
	"github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain"
	"time"
)

type Repository interface {
	GetDay(int64, time.Time) (*domain.Day, error)
	CreateDay(*domain.Day) error
	UpdateDay(*domain.Day) error
}
