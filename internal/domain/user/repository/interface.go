package user_repo

import "github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain/user"

type UserRepository interface {
	Create(user *user.User) error
	Get(id int64) (*user.User, error)
	Update(user user.User) error
}
