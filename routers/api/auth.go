package api

import (
	"flash_sale/pkg/app"
	"flash_sale/pkg/constants"
	"flash_sale/pkg/util"
	"flash_sale/service/auth_service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Summary Get Auth
// @Produce  json
// @Tags Auth
// @Param auth body app.Auth true "Auth request body"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	var (
		appG = app.Gin{C: c}
		data app.Auth
	)

	httpCode, errCode := app.BindAndValid(c, &data)
	if errCode != constants.SUCCESS {
		appG.Response(httpCode, errCode, nil)
		return
	}

	authService := auth_service.Auth{Username: data.Username, Password: data.Password}
	isExist, err := authService.Check()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}

	if !isExist {
		appG.Response(http.StatusUnauthorized, constants.ERROR_AUTH, nil)
		return
	}
	userID, err := authService.GetUserID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_AUTH_CHECK_TOKEN_FAIL, nil)
		return
	}
	token, err := util.GenerateToken(data.Username, data.Password, userID)
	if err != nil {
		appG.Response(http.StatusInternalServerError, constants.ERROR_AUTH_TOKEN, nil)
		return
	}

	appG.Response(http.StatusOK, constants.SUCCESS, map[string]string{
		"token": token,
	})
}
