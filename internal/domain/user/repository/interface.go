package user_repo

import "github.com/kaverhovsky/gosniias-time-manager-bot/internal/domain/user"

type UserRepository interface {
	Create(user *user.User)
	Get(id int64)
	Update(user user.User)
}
