package day_item

import (
	"github.com/google/uuid"
	"time"
)

type DayRecordItem struct {
	ID       uuid.UUID
	Type     string
	Datetime time.Time
}
