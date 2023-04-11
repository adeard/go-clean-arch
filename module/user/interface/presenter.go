package user

import "go-clean-arch/model"

type Presenter interface {
	ResponseUsers(us []*model.User) []*model.User
}
