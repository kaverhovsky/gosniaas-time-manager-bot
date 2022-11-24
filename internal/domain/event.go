package domain

import "time"

type EventType int

type Event struct {
	Type     string
	Datetime time.Time
}
