package scheduler

import (
	"encoding/csv"
	"errors"
	"github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain"
	"os"
	"time"
)

const (
	// TODO добавить layout
	layout = ""
)

func durationsFromCSV(filename string) (map[int64]time.Duration, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	durations := make(map[int64]time.Duration)

	for _, row := range rows {
		if len(row) != 2 {
			return nil, errors.New("wrong row format in \"durations\" file")
		}

		datetime, err := time.Parse(layout, row[0])
		if err != nil {
			// TODO обернуть ошибку
			return nil, err
		}

		// add record to scheduler's durations
		dur, err := time.ParseDuration(row[1])
		if err != nil {
			// TODO обернуть ошибку
			return nil, err
		}
		durations[datetime.Unix()] = dur
	}

	return durations, nil
}

func periodsFromCSV(filename string) ([]*domain.Period, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	r := csv.NewReader(f)
	rows, err := r.ReadAll()
	if err != nil {
		return nil, err
	}

	var periods []*domain.Period

	for _, row := range rows {
		if len(row) != 2 {
			return nil, errors.New("wrong row format in \"periods\" file")
		}

		start, err := time.Parse(layout, row[0])
		if err != nil {
			// TODO обернуть ошибку
			return nil, err
		}
		end, err := time.Parse(layout, row[0])
		if err != nil {
			// TODO обернуть ошибку
			return nil, err
		}

		periods = append(periods, &domain.Period{Start: start, End: end})
	}

	return periods, nil
}
