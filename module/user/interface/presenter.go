package user

import "go-clean-arch/domain/model"

type Presenter interface {
	ResponseUsers(us []*model.User) []*model.User
}
