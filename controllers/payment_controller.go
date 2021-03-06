package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucaspizzo/api-payment/forms"
	"github.com/lucaspizzo/api-payment/services"
	"github.com/pkg/errors"
)

type Paymenter interface {
	AddPayment(ctx *gin.Context)
}

type PaymentController struct {
	BaseController
	PaymentService services.PaymentContract `inject:""`
}

func (p *PaymentController) AddPayment(ctx *gin.Context) {
	form := make([]*forms.PaymentForm, 0)
	ctx.ShouldBind(&form)

	if len(form) == 0 {
		p.respondError(ctx, errors.New("Empty body"))
		return
	}

	if !forms.ValidatePaymentFormList(&form) {
		p.respond(ctx, form)
		return
	}

	err := p.PaymentService.RegisterPayments(&form)

	if err != nil {
		p.respondError(ctx, err)
		return
	}

	p.respondSuccessNoContent(ctx)
}
