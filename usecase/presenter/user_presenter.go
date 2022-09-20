package presenter

import "go-clean-arch/domain/model"

type UserPresenter interface {
	ResponseUsers(u []*model.User) []*model.User
}
