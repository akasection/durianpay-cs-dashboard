package jwt

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/common"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/util"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var data interface{}

		code = common.SUCCESS_OK

		// Token retrieval from query string or header
		token := c.Query("token")
		headerToken := c.Request.Header.Get("Authorization")
		prefix := "Bearer "
		if len(headerToken) > len(prefix) && headerToken[:len(prefix)] == prefix {
			token = headerToken[len(prefix):]
		}

		if token == "" {
			code = common.ERROR_USER_MISSING_TOKEN
		} else {
			_, err := util.ParseToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					code = common.ERROR_USER_TOKEN_EXPIRED
				default:
					code = common.ERROR_USER_TOKEN_INVALID
				}
			}
		}
		if code != common.SUCCESS_OK {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  common.MessageCode(code),
				"data": data,
			})

			c.Abort()
			return
		}

		c.Next()
	}
}
