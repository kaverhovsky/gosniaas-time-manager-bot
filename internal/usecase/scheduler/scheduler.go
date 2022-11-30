package scheduler

import (
	"github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain"
	"go.uber.org/zap"
	"time"
)

const (
	durationsFilename = "durations.csv"
	periodsFilename   = "periods.csv"
)

type scheduler struct {
	durations map[int64]time.Duration
	periods   []*domain.Period
	logger    *zap.Logger
}

var sch *scheduler

func Get() Scheduler {
	if sch == nil {
		return nil
	}
	return sch
}

func New(logger *zap.Logger) *scheduler {
	return &scheduler{
		durations: make(map[int64]time.Duration),
		logger:    logger,
	}
}

func (sch *scheduler) init() error {
	// TODO добавить логи
	durations, err := durationsFromCSV(durationsFilename)
	if err != nil {
		return err
	}
	periods, err := periodsFromCSV(durationsFilename)
	if err != nil {
		return err
	}
	sch.durations = durations
	sch.periods = periods

	return nil
}

func (sch *scheduler) GetPeriod(now time.Time) *domain.Period {
	for _, period := range sch.periods {
		if period.Start.Before(now) && period.End.After(now) {
			return period
		}
	}
	return nil
}

func (sch *scheduler) SumForNow(now time.Time) (time.Duration, error) {
	//period := sch.GetPeriod(now)
	//ts := period.Start.Truncate()
	//for ts = period.Start; !
	return 0, nil
}
