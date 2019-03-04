package controllers

import (
	"github.com/lucaspizzo/api-payment/forms"
	"github.com/lucaspizzo/api-payment/services/exceptions"
	"net/http"
	"reflect"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

func (b *BaseController) GetStoreID(c *gin.Context) uint {
	return c.MustGet("storeID").(uint)
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
	}

	ctx.JSON(http.StatusOK, result)
}

func (b *BaseController) respondSucessNoContent(ctx *gin.Context) {
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