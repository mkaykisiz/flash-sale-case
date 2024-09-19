package jwt

import (
	"flash_sale/pkg/constants"
	"flash_sale/pkg/util"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = constants.SUCCESS
		token := c.GetHeader("token")
		if token == "" {
			code = constants.INVALIDPARAMS
		} else {
			claims, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = constants.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				default:
					code = constants.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			} else {
				c.Set("user_id", claims.UserID)
			}
		}

		if code != constants.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  constants.GetMsg(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
