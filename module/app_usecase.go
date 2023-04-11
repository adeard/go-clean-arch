package module

import user "go-clean-arch/module/user/interface"

type AppUsecase interface {
	user.Usecase
}
