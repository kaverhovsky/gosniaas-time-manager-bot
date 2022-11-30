package day_repo

import (
	"github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain/day"
	"time"
)

type DayRepository interface {
	Get(uID int64, year int, month string, day int) (*day.DayRecord, error)
	GetMany(uID int64, start time.Time, end time.Time) ([]*day.DayRecord, error)
	Create(*day.DayRecord) error
	Update(*day.DayRecord) error
}
