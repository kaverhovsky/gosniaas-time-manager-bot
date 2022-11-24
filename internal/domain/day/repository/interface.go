package repository

import (
	"github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain/day"
)

type Repository interface {
	Get(UID int64, year int, month string, day int) (*day.DayRecord, error)
	Create(*day.DayRecord) error
	Update(*day.DayRecord) error
}
