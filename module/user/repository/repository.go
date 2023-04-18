package user

import (
	"go-clean-arch/model"
	user "go-clean-arch/module/user/interface"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{db}
}

func (ur *userRepository) FindAll(u []*model.User) ([]*model.User, error) {
	err := ur.db.Find(&u).Error
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (ur *userRepository) GetUserByEmail(u []*model.User, email string) ([]*model.User, error) {
	result := ur.db.Where("email = ?", email).First(&u)

	if result.Error != nil {
		return nil, result.Error
	}

	return u, nil
}

func (ur *userRepository) GetUserById(u *model.User, id int) (*model.User, error) {

	u = &model.User{}

	result := ur.db.Where("id = ?", id).First(&u)

	if result.Error != nil {
		return nil, result.Error
	}
	return u, nil
}

func (ur *userRepository) UpdateUserById(u *model.User, id int, c echo.Context) (*model.User, error) {
	user := new(model.User)
	user.Nama = c.FormValue("nama")
	user.Email = c.FormValue("email")

	result := ur.db.Model(&model.User{}).Where("id = ?", id).Updates(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	updated, error := ur.GetUserById(u, id)

	if error != nil {
		return nil, error
	}

	return updated, nil
}
