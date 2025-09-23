package AuthApi

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/akasection/durianpay-cs-dashboard/backend/models"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/common"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/ginutil"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/util"
)

type loginData struct {
	Username string `json:"username" binding:"required" validate:"min=3,max=31,required"`
	Password string `json:"password" binding:"required" validate:"min=4,max=32,required"`
}

func PostLogin(c *gin.Context) {
	appG := ginutil.Gin{C: c}
	var data loginData
	if err := c.BindJSON(&data); err != nil {
		appG.SendResponse(http.StatusBadRequest, common.ERROR_USER_CREDENTIALS_INVALID, nil, nil)
		return
	}

	result := models.CheckCredentials(data.Username, data.Password)

	if !result {
		appG.SendResponse(http.StatusUnauthorized, common.ERROR_USER_CREDENTIALS_INVALID, nil, nil)
	}

	token, tokenErr := util.GenerateToken(data.Username, data.Password)
	if tokenErr != nil {
		appG.SendResponse(http.StatusInternalServerError, common.ERROR_GENERIC, nil, nil)
		return
	}

	appG.SendResponse(http.StatusOK, common.SUCCESS_OK, map[string]string{
		"token": token,
	}, nil)

}
