package user

import "go-clean-arch/model"

type Presenter interface {
	ResponseUsers(us []*model.User) []*model.User
	ResponseUsersList(us []*model.User) *model.UserList
	ResponseUsersDetail(us []*model.User) *model.UserDetail
}
