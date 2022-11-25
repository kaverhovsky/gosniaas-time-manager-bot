package usecase

import (
	"errors"
	"github.com/google/uuid"
	"github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain"
	"github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain/day"
	day_repo "github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain/day/repository"
	"github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain/day_item"
	day_item_repo "github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain/day_item/repository"
	user_repo "github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain/user/repository"
	"github.com/kaverhovsky/gosniias-time-manager-bot/internal/usecase/scheduler"
	"github.com/kaverhovsky/gosniias-time-manager-bot/pkg/common"
	"go.uber.org/zap"
	"time"
)

type Usecase struct {
	logger      *zap.Logger
	config      *common.Config
	dayRepo     day_repo.DayRepository
	dayItemRepo day_item_repo.DayItemRepository
	userRepo    user_repo.UserRepository
}

func New(logger *zap.Logger,
	config *common.Config,
	dayRepo day_repo.DayRepository,
	dayItemRepo day_item_repo.DayItemRepository,
	userRepo user_repo.UserRepository) *Usecase {
	return &Usecase{
		logger:      logger,
		config:      config,
		dayRepo:     dayRepo,
		dayItemRepo: dayItemRepo,
		userRepo:    userRepo,
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

	day, err := u.dayRepo.Get(event.UID, y, m.String(), d)
	// TODO обрабатывать repo ошибку NotFound
	if err != nil {
		u.logger.Error("failed getting day")
		return err
	}

	if day == nil {
		id, err := u.createDay(event)
		if err != nil {
			u.logger.Error("failed creating day")
			return err
		}
		if err = u.createDayItem(id, event); err != nil {
			u.logger.Error("failed getting day item")
			return err
		}
		return nil
	}

	//if day.LastOutTime.IsZero() || day.LastOutTime.After(event.Datetime) {
	//	// TODO возвращать ошибки с контекстом (инстансом), о котором идет речь
	//	return errors.New("user had entered already that day")
	//}

	if err = u.UpdateDay(day, event); err != nil {
		u.logger.Error("failed updating day")
		return err
	}
	if err = u.createDayItem(day.ID, event); err != nil {
		u.logger.Error("failed getting day item")
		return err
	}
	return nil
}

func (u *Usecase) createDay(event *domain.Event) (uuid.UUID, error) {
	y, m, d := event.Datetime.Date()
	id := uuid.New()
	day := &day.DayRecord{
		ID:    id,
		UID:   event.UID,
		Year:  y,
		Month: m.String(),
		Day:   d,
		//SumHours:    0,
		FirstInTime: time.Now().UTC(),
		LastOutTime: time.Time{},
	}

	err := u.dayRepo.Create(day)
	if err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (u *Usecase) UpdateDay(day *day.DayRecord, event *domain.Event) error {
	if event.Type == domain.Entrance.String() {
		if day.LastInTime.After(day.LastOutTime) {
			return errors.New("double entrance occurred")
		}
		day.LastInTime = event.Datetime
	} else {
		if day.LastOutTime.After(day.LastInTime) {
			return errors.New("double entrance occurred")
		}
		day.LastOutTime = event.Datetime
	}
	if err := u.dayRepo.Update(day); err != nil {
		return err
	}
	return nil
}

func (u *Usecase) createDayItem(dayID uuid.UUID, event *domain.Event) error {
	item := &day_item.DayRecordItem{
		ID:       uuid.New(),
		DayID:    dayID,
		Type:     event.Type,
		Datetime: event.Datetime,
	}

	if err := u.dayItemRepo.Create(item); err != nil {
		return err
	}
	return nil
}

func (u *Usecase) Report(uid int64) (*domain.Report, error) {
	user, err := u.userRepo.Get(uid)
	if err != nil {
		return nil, err
	}
	now := time.Now()
	year, month, _ := now.Date()
	days, err := u.dayRepo.GetMany(user.ID, year, month.String())
	if err != nil {
		return nil, err
	}

	var sumHours time.Duration
	for _, day := range days {
		sumHoursForDay := day.LastOutTime.Sub(day.FirstInTime)
		sumHours += sumHoursForDay
	}

	idealSumHours := scheduler.Get().GetSum(year, month)
	diffHours := sumHours - idealSumHours
	return &domain.Report{
		Firstname:     user.Firstname,
		Lastname:      user.Lastname,
		Year:          year,
		Month:         month.String(),
		SumHours:      sumHours,
		IdealSumHours: idealSumHours,
		DiffHours:     diffHours,
	}, nil
}
