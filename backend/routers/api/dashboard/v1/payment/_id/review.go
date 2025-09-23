package PaymentReviewApi

import (
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/common"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/ginutil"
	"github.com/gin-gonic/gin"
)

type ReviewAction struct {
	PaymentId string `json:"payment_id" binding:"required"`
	Action    string `json:"action" binding:"required,oneof=approve reject"`
}

func PutReviewPayment(c *gin.Context) {
	appG := ginutil.Gin{C: c}
	paymentId := c.Param("id")
	var reviewAction = ReviewAction{
		PaymentId: paymentId,
		Action:    c.PostForm("action"),
	}

	if err := c.ShouldBindJSON(&reviewAction); err != nil {
		appG.SendResponse(400, common.ERROR_INVALID_PARAMS, nil)
		return
	}

	appG.SendResponse(200, common.SUCCESS_OK, map[string]string{
		"payment_id": reviewAction.PaymentId,
		"action":     reviewAction.Action,
	})
}
