package controller

import user "go-clean-arch/module/user/interface"

type AppController interface {
	user.Usecase
}
