package user

import (
	"go-clean-arch/model"
	user "go-clean-arch/module/user/interface"
)

type userPresenter struct {
}

func NewUserPresenter() user.Presenter {
	return &userPresenter{}
}

func (up *userPresenter) ResponseUsers(us []*model.User) []*model.User {
	for _, u := range us {
		u.Nama = "Mr." + u.Nama
	}
	return us
}
