package repository

import "go-clean-arch/domain/model"

type UserRepository interface {
	FindAll(u []*model.User) ([]*model.User, error)
}
