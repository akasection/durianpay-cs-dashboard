package middleware

import (
	"net/http"

	"github.com/akasection/durianpay-cs-dashboard/backend/models"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/common"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/ginutil"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/util"
	"github.com/gin-gonic/gin"
)

func UseRbac(roles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := ginutil.Gin{C: c}
		token, _ := util.GetTokenFromRequest(c)
		claim, _ := util.ParseToken(token)

		userRoles, _ := models.GetUserRoles(claim.Username)
		if util.Intersection(userRoles, roles) == nil {
			appG.SendResponse(http.StatusForbidden, common.ERROR_INSUFFICIENT_PERMISSIONS, nil, nil)
			return
		}

		c.Next()
	}

}
