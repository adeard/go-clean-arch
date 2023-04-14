package user

import (
	"go-clean-arch/model"
	user "go-clean-arch/module/user/interface"

	"github.com/labstack/echo/v4"
)

type userHandlers struct {
	UserRepository user.Repository
	UserPresenter  user.Presenter
}

func NewUserHandlers(r user.Repository, p user.Presenter) user.Handler {
	return &userHandlers{r, p}
}

func (us *userHandlers) Get(u []*model.User) (*model.UserList, error) {
	u, err := us.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsersList(u), nil
}

func (us *userHandlers) GetUserByEmail(u []*model.User, email string) ([]*model.User, error) {
	u, err := us.UserRepository.GetUserByEmail(u, email)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}

func (us *userHandlers) GetUserById(u []*model.User, id int) (*model.UserDetail, error) {
	u, err := us.UserRepository.GetUserById(u, id)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsersDetail(u), nil
}

func (us *userHandlers) UpdateUserById(u []*model.User, id int, c echo.Context) (*model.UserDetail, error) {
	update, err := us.UserRepository.UpdateUserById(u, id, c)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsersDetail(update), nil
}
