package repository

import "github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain/day"

type Repository interface {
	Get(UID int64, year int, month string, day int) (*day.DayRecord, error)
	CreateDay(*day.DayRecord) error
	UpdateDay(*day.DayRecord) error
}
