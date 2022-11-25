package domain

import "time"

type EventType string

func (et EventType) String() string {
	return string(et)
}

const (
	Entrance EventType = "in"
	Exit               = "out"
)

type Event struct {
	UID      int64
	Type     string
	Datetime time.Time
}
