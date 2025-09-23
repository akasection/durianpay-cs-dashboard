package middleware

import (
	"log"
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
		var userRolesStr []string = util.Intersection(userRoles, roles)
		log.Println("roles", roles, "userRoles", userRoles, "userRolesStr:", userRolesStr)
		if len(userRolesStr) == 0 {
			appG.SendResponse(http.StatusForbidden, common.ERROR_MISMATCHED_ROLE, nil, nil)
			c.Abort()
			return
		}

		c.Next()
	}

}
