package main

import (
	"fmt"
	"github.com/facebookgo/inject"
	_ "github.com/joho/godotenv/autoload"
	"github.com/lucaspizzo/api-payment/api"
	"github.com/lucaspizzo/api-payment/controllers"
	db "github.com/lucaspizzo/api-payment/infrastructure/database"
	"github.com/lucaspizzo/api-payment/infrastructure/repositories"
	"github.com/lucaspizzo/api-payment/services"
	"os"
)

func main() {
	BuildContainer()
}


func BuildContainer() {
	var g inject.Graph
	db := &db.Repository{}
	// httpClient := &http.Client{}
	db.Start()
	gorm := db.GetInstance()

	server := &api.Server{}

	accountRepository := &repositories.AccountRepository{DB: gorm}

	accountService := &services.AccountService{}

	accountController := &controllers.AccountController{}


	err := g.Provide(
		&inject.Object{Value: server},
		&inject.Object{Value: accountRepository},
		&inject.Object{Value: accountService},
		&inject.Object{Value: accountController},
	)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err := g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	server.Run()
}