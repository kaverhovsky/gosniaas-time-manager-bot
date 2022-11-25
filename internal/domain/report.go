package domain

import "time"

type Report struct {
	Firstname     string
	Lastname      string
	Year          int
	Month         string
	SumHours      time.Duration
	IdealSumHours time.Duration
	DiffHours     time.Duration
}
