package interactor

import (
	"go-clean-arch/domain/model"
	"go-clean-arch/usecase/presenter"
	"go-clean-arch/usecase/repository"

	"github.com/labstack/echo"
)

type userInteractor struct {
	UserRepository repository.UserRepository
	UserPresenter  presenter.UserPresenter
}

type UserInteractor interface {
	Get(u []*model.User) ([]*model.User, error)
	GetUserByEmail(u []*model.User, email string) ([]*model.User, error)
	GetUserById(u []*model.User, id int) ([]*model.User, error)
	UpdateUserById(u []*model.User, id int, c echo.Context) ([]*model.User, error)
}

func NewUserInteractor(r repository.UserRepository, p presenter.UserPresenter) UserInteractor {
	return &userInteractor{r, p}
}

func (us *userInteractor) Get(u []*model.User) ([]*model.User, error) {
	u, err := us.UserRepository.FindAll(u)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}

func (us *userInteractor) GetUserByEmail(u []*model.User, email string) ([]*model.User, error) {
	u, err := us.UserRepository.GetUserByEmail(u, email)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}

func (us *userInteractor) GetUserById(u []*model.User, id int) ([]*model.User, error) {
	u, err := us.UserRepository.GetUserById(u, id)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}

func (us *userInteractor) UpdateUserById(u []*model.User, id int, c echo.Context) ([]*model.User, error) {
	u, err := us.UserRepository.UpdateUserById(u, id, c)
	if err != nil {
		return nil, err
	}

	return us.UserPresenter.ResponseUsers(u), nil
}
