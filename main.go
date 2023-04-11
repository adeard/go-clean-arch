package main

import (
	"fmt"
	"go-clean-arch/config"
	"go-clean-arch/datastore"
	user "go-clean-arch/module/user/delivery"
	"go-clean-arch/registry"
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = user.UserRoutes(e, r.NewAppUsecase())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
