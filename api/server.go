package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lucaspizzo/api-payment/controllers"
	"github.com/lucaspizzo/api-payment/infrastructure/config"
	"github.com/lucaspizzo/api-payment/infrastructure/database/migrations"
)

type Server struct {
	AccountController     controllers.Accounter     `inject:""`
	TransactionController controllers.Transactioner `inject:""`
	PaymentController     controllers.Paymenter     `inject:""`
}

func (s *Server) Run() {
	migrations.RunMigrations()
	r := s.SetupRoutes()
	r.Run(":" + config.APP_PORT)
}

func (s *Server) SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/health/check", func(ctx *gin.Context) {
		ctx.JSON(200, "OK")
	})

	v1 := router.Group("/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET("/limits", s.AccountController.ListAccount)
			accounts.PATCH("/:id", s.AccountController.UpdateLimits)
		}

		v1.POST("/transactions", s.TransactionController.AddTransaction)
		v1.POST("/payments", s.PaymentController.AddPayment)
	}

	return router
}
