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

	// DB
	db := &db.Repository{}
	db.Start()
	gorm := db.GetInstance()

	// CORE
	server := &api.Server{}

	// REPOSITORIES
	accountRepository := &repositories.AccountRepository{DB: gorm}
	operationTypeRepository := &repositories.OperationTypeRepository{DB: gorm}
	transactionRepository := &repositories.TransactionRepository{DB: gorm}

	// SERVICES
	accountService := &services.AccountService{}
	operationTypeService := &services.OperationTypeService{}
	transactionService := &services.TransactionService{}
	paymentService := &services.PaymentService{}

	// CONTROLLERS
	accountController := &controllers.AccountController{}
	transactionController := &controllers.TransactionController{}
	paymentController := &controllers.PaymentController{}

	err := g.Provide(

		// DB
		&inject.Object{Value: db},

		// CORE
		&inject.Object{Value: server},

		// REPOSITORIES
		&inject.Object{Value: accountRepository},
		&inject.Object{Value: operationTypeRepository},
		&inject.Object{Value: transactionRepository},

		// SERVICES
		&inject.Object{Value: accountService},
		&inject.Object{Value: operationTypeService},
		&inject.Object{Value: transactionService},
		&inject.Object{Value: paymentService},

		// CONTROLLERS
		&inject.Object{Value: accountController},
		&inject.Object{Value: transactionController},
		&inject.Object{Value: paymentController},
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
