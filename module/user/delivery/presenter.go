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

func (up *userPresenter) ResponseUsersList(us []*model.User) *model.UserList {

	result := &model.UserList{
		Status:   200,
		ErrorMsg: " ",
		Data:     us,
	}

	return result
}

func (up *userPresenter) ResponseUsersDetail(us []*model.User) *model.UserDetail {

	result := &model.UserDetail{
		Status:   200,
		ErrorMsg: " ",
		Data:     us[0],
	}

	return result
}
