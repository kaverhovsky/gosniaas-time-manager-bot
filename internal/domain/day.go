package domain

import "time"

type Day struct {
	UID         int64
	Date        time.Time
	SumHours    float32
	FirstInTime time.Time
	LastOutTime time.Time
}
