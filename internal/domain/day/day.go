package day

import (
	"github.com/google/uuid"
	"time"
)

type DayRecord struct {
	ID    uuid.UUID
	UID   int64
	Year  int
	Month string
	Day   int
	//SumHours    time.Duration
	FirstInTime time.Time
	LastInTime  time.Time
	LastOutTime time.Time
}

//type DayRecordUpdate struct {
//	UID         *int64
//	Year        *int
//	Month       *string
//	Day         *int
//	SumHours    *float32
//	FirstInTime time.Time
//	LastOutTime time.Time
//}
