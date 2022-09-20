package registry

import (
	"go-clean-arch/interface/controller"

	ip "go-clean-arch/interface/presenters"

	ir "go-clean-arch/interface/repository"

	"go-clean-arch/usecase/interactor"

	up "go-clean-arch/usecase/presenter"

	ur "go-clean-arch/usecase/repository"
)

func (r *registry) NewUserController() controller.UserController {
	return controller.NewUserController(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() interactor.UserInteractor {
	return interactor.NewUserInteractor(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() ur.UserRepository {
	return ir.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() up.UserPresenter {
	return ip.NewUserPresenter()
}
