package registry

import (
	"go-clean-arch/module"

	"github.com/jinzhu/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppUsecase() module.AppUsecase
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppUsecase() module.AppUsecase {
	return r.NewUserUsecase()
}
