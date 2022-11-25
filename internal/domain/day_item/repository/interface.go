package day_item_repo

import (
	"github.com/google/uuid"
	"github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain/day_item"
)

type DayItemRepository interface {
	Get(id uuid.UUID) (*day_item.DayRecordItem, error)
	GetMany(dayID uuid.UUID) ([]*day_item.DayRecordItem, error)
	Create(item *day_item.DayRecordItem) error
}
