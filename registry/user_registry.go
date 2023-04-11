package registry

import (
	userUsecase "go-clean-arch/module/user/usecase"

	userInterface "go-clean-arch/module/user/interface"

	userDelivery "go-clean-arch/module/user/delivery"

	userRepository "go-clean-arch/module/user/repository"
)

func (r *registry) NewUserUsecase() userInterface.Usecase {
	return userUsecase.NewUserUsecase(r.NewUserInteractor())
}

func (r *registry) NewUserInteractor() userInterface.Handler {
	return userDelivery.NewUserHandlers(r.NewUserRepository(), r.NewUserPresenter())
}

func (r *registry) NewUserRepository() userInterface.Repository {
	return userRepository.NewUserRepository(r.db)
}

func (r *registry) NewUserPresenter() userInterface.Presenter {
	return userDelivery.NewUserPresenter()
}
