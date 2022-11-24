package repository

import (
	"github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain"
)

type Repository interface {
	GetDay(UID int64, year int, month string, day int) (*domain.DayRecord, error)
	CreateDay(*domain.DayRecord) error
	UpdateDay(*domain.DayRecord) error
}
