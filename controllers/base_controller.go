package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lucaspizzo/api-payment/forms"
	"github.com/lucaspizzo/api-payment/services/exceptions"
	"net/http"
	"reflect"
)

type BaseController struct {
}

func (b *BaseController) respond(ctx *gin.Context, result interface{}) {
	value := reflect.ValueOf(result)

	if value.Kind() == reflect.Struct {
		field := value.FieldByName("BaseForm")

		if field.IsValid() {
			value := field.Interface()
			form, ok := value.(forms.BaseForm)

			if ok && !form.IsValid() {
				ctx.JSON(http.StatusUnprocessableEntity, form.GetErrors())
				return
			}
		}
	} else if value.Kind() == reflect.Slice {
		if values, ok := result.([]*forms.PaymentForm); ok {
			for _, v := range values {
				if !v.IsValid() {
					ctx.JSON(http.StatusUnprocessableEntity, v.GetErrors())
					return
				}
			}
		}
	}

	ctx.JSON(http.StatusOK, result)
}

func (b *BaseController) respondSuccessNoContent(ctx *gin.Context) {
	ctx.Status(http.StatusNoContent)
}

func (b *BaseController) respondError(ctx *gin.Context, err error) {
	switch err.(type) {
	case *exceptions.InvalidEntityError:
		ctx.JSON(http.StatusUnprocessableEntity, err.Error())
		break
	case *exceptions.NotFoundEntityError:
		ctx.JSON(http.StatusNotFound, err.Error())
		break

	default:
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

}
