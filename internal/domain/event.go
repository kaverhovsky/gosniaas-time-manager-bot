package domain

import "time"

type EventType int

type Event struct {
	UID      int64
	Type     string
	Datetime time.Time
}
