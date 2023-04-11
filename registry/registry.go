package registry

import (
	"go-clean-arch/interface/controller"

	"github.com/jinzhu/gorm"
)

type registry struct {
	db *gorm.DB
}

type Registry interface {
	NewAppUsecase() controller.AppController
}

func NewRegistry(db *gorm.DB) Registry {
	return &registry{db}
}

func (r *registry) NewAppUsecase() controller.AppController {
	return r.NewUserController()
}
