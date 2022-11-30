package scheduler

import (
	"encoding/csv"
	"errors"
	"github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain"
	"go.uber.org/zap"
	"os"
	"time"
)

const (
	durationsFilename = "durations.csv"
	layout            = ""
)

type scheduler struct {
	durations map[int64]time.Duration
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
	f, err := os.Open(durationsFilename)
	if err != nil {
		return err
	}
	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		return err
	}
	for _, row := range rows {
		if len(row) != 2 {
			return errors.New("wrong row format in \"durations\" file")
		}

		datetime, err := time.Parse(layout, row[0])
		if err != nil {
			// TODO обернуть ошибку
			return err
		}

		// add record to scheduler's durations
		dur, err := time.ParseDuration(row[1])
		if err != nil {
			// TODO обернуть ошибку
			return err
		}
		sch.durations[datetime.Unix()] = dur
	}

	return nil
}

func (sch *scheduler) GetPeriod(now time.Time) *domain.Period {
	//TODO implement me
	panic("implement me")
}

func (sch *scheduler) SumForNow(now time.Time) time.Duration {
	//TODO implement me
	panic("implement me")
}
