package app

import (
	"flash_sale/pkg/constants"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

// BindAndValid binds and validates data
func BindAndValid(c *gin.Context, data interface{}) (int, int) {
	err := c.Bind(data)
	if err != nil {
		return http.StatusBadRequest, constants.INVALIDPARAMS
	}

	valid := validation.Validation{}
	check, err := valid.Valid(data)
	if err != nil {
		return http.StatusInternalServerError, constants.ERROR
	}
	if !check {
		MarkErrors(valid.Errors)
		return http.StatusBadRequest, constants.INVALIDPARAMS
	}

	return http.StatusOK, constants.SUCCESS
}
