package controllers

import (

	"github.com/gin-gonic/gin"
	"github.com/lucaspizzo/api-payment/domains/renders"
	"github.com/lucaspizzo/api-payment/forms"
	"github.com/lucaspizzo/api-payment/services"
	"strconv"
)

type Accounter interface {
	RegisterAccount(ctx *gin.Context)
	RetrieveAccount(ctx *gin.Context)
	ListAccount(ctx *gin.Context)
	UpdateLimits(ctx *gin.Context)
}

type AccountController struct {
	BaseController
	AccountService services.AccountContract `inject:""`
}

func (a *AccountController) RegisterAccount(ctx *gin.Context) {
	a.respond(ctx, renders.NewPaymentTypeResponse(nil))
}

func (a *AccountController) RetrieveAccount(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)


	account, err := a.AccountService.Get(id)
	if err != nil {
		a.respondError(ctx, err)
		return
	}
	a.respond(ctx, renders.NewPaymentTypeResponse(account))
}

func (a *AccountController) ListAccount(ctx *gin.Context) {
	accounts, err := a.AccountService.List()

	if err != nil {
		a.respondError(ctx, err)
		return
	}
	a.respond(ctx, renders.NewPaymentTypeResponseList(accounts))
}

func (a *AccountController) UpdateLimits(ctx *gin.Context) {
	id, _ := strconv.ParseUint(ctx.Param("id"), 10, 64)

	form := forms.LimitForm{AccountID: id}
	ctx.ShouldBind(&form)

	if !form.Validate() {
		a.respond(ctx, form)
		return
	}


	account, err := a.AccountService.UpdateLimits(&form)

	if err != nil {
		a.respondError(ctx, err)
		return
	}
	a.respond(ctx, renders.NewPaymentTypeResponse(account))
}
