package api

import (
	"github.com/gin-gonic/gin"
	"github.com/lucaspizzo/api-payment/controllers"
	"github.com/lucaspizzo/api-payment/infrastructure/config"
	"github.com/lucaspizzo/api-payment/infrastructure/database/migrations"
)

type Server struct {
	AccountController controllers.Accounter `inject:""`
}

func (s *Server) Run() {
	migrations.RunMigrations()
	r := s.SetupRoutes()
	r.Run(":" + config.APP_PORT)
}

func (s *Server) SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/health/check", func (ctx *gin.Context) {
		ctx.JSON(200, "OK")
	})

	router.GET("/v1/accounts/limits", s.AccountController.ListAccount)
	router.PATCH("/v1/accounts/:id", s.AccountController.UpdateLimits)

	return router
}