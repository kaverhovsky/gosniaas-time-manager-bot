package domain

import (
	"github.com/google/uuid"
	"time"
)

type DayRecord struct {
	ID          uuid.UUID
	UID         *int64
	Year        *int
	Month       *string
	Day         *int
	SumHours    *float32
	FirstInTime *time.Time
	LastOutTime *time.Time
}

type DayRecordItem struct {
	ID       uuid.UUID
	Type     string
	Datetime time.Time
}
