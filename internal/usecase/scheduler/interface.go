package scheduler

import (
	"github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain"
	"time"
)

type Scheduler interface {
	GetPeriod(now time.Time) *domain.Period
	SumForNow(now time.Time) (time.Duration, error)
}
