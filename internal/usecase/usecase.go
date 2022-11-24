package usecase

import (
	"errors"
	"github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain"
	"github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain/day"
	"github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain/day/repository"
	"github.com/kaverhovsky/gosniias-time-manager-bot/pkg/common"
	"go.uber.org/zap"
	"time"
)

type Usecase struct {
	logger *zap.Logger
	config *common.Config
	repo   repository.Repository
}

func New(logger *zap.Logger, config *common.Config, repo repository.Repository) *Usecase {
	return &Usecase{
		logger: logger,
		config: config,
		repo:   repo,
	}
}

func (u *Usecase) SaveEntrance(event *domain.Event) error {
	// TODO ввести dto для верхнего уровня, в котором будет полный time.Time
	y, m, d := event.Datetime.Date()
	//const dateLayout = "2006-01-02"
	//date, err := time.Parse(dateLayout, fmt.Sprintf("%d-%d-%d", y, m, d))
	//if err != nil {
	//	return errors.New("can't parse date from event")
	//}

	day, err := u.repo.GetDay(event.UID, y, m.String(), d)
	// TODO обрабатывать repo ошибку NotFound
	if err != nil {
		return errors.New("failed getting day for UID and Date")
	}

	if day == nil {
		if err := u.createDay(event); err != nil {
			return err
		}
		return nil
	}

	if day.LastOutTime.IsZero() || day.LastOutTime.After(event.Datetime) {
		// TODO возвращать ошибки с контекстом (инстансом), о котором идет речь
		return errors.New("user had entered already that day")
	}

}

func (u *Usecase) createDay(event *domain.Event) error {
	y, m, d := event.Datetime.Date()
	day := &day.DayRecord{
		UID:         event.UID,
		Year:        y,
		Month:       m.String(),
		Day:         d,
		SumHours:    0,
		FirstInTime: time.Now().UTC(),
		LastOutTime: time.Time{},
	}

	err := u.repo.CreateDay(day)
	if err != nil {
		return err
	}

	u.repo.CreateDayItem()

	return nil
}
