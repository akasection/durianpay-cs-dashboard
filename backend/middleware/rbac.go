package middleware

import (
	"net/http"

	"github.com/akasection/durianpay-cs-dashboard/backend/models"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/common"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/ginutil"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/util"
	"github.com/gin-gonic/gin"
)

func UseRbac(permissions ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		appG := ginutil.Gin{C: c}
		token, _ := util.GetTokenFromRequest(c)
		claim, _ := util.ParseToken(token)

		userPermission, _ := models.GetUserPermissions(claim.Username)
		if util.Intersection(userPermission, permissions) == nil {
			appG.SendResponse(http.StatusForbidden, common.ERROR_INSUFFICIENT_PERMISSIONS, nil, nil)
			return
		}

		c.Next()
	}

}
