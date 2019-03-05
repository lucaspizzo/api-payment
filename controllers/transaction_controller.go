package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucaspizzo/api-payment/forms"
	"github.com/lucaspizzo/api-payment/services"
)

type Transactioner interface {
	AddTransaction(ctx *gin.Context)
}

type TransactionController struct {
	BaseController
	TransactionService services.TransactionContract `inject:""`
}

func (t *TransactionController) AddTransaction(ctx *gin.Context) {
	form := forms.TransactionForm{}
	ctx.ShouldBind(&form)

	if !form.Validate() {
		t.respond(ctx, form)
		return
	}

	transaction, err := t.TransactionService.RegisterTransaction(&form)

	if err != nil {
		t.respondError(ctx, err)
		return
	}

	t.respond(ctx, transaction)
}
