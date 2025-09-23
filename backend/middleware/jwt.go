package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"

	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/common"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/ginutil"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/util"
)

func UseJwt() gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := ginutil.Gin{C: c}
		var code int

		code = common.SUCCESS_OK

		// Token retrieval from query string or header
		token, err := util.GetTokenFromRequest(c)

		if token == "" || err != nil {
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
			appG.SendResponse(http.StatusUnauthorized, code, nil, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
