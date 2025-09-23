package PaymentReviewApi

import (
	"net/http"
	"strconv"

	"github.com/akasection/durianpay-cs-dashboard/backend/models"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/common"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/ginutil"
	"github.com/gin-gonic/gin"
)

type ReviewAction struct {
	PaymentId uint                 `json:"payment_id" binding:"required"`
	Action    models.PaymentStatus `json:"action" binding:"required,oneof=completed failed"`
}

func PutReviewPayment(c *gin.Context) {
	appG := ginutil.Gin{C: c}
	paymentId, _ := strconv.Atoi(c.Param("id"))
	var payload = ReviewAction{
		PaymentId: uint(paymentId),
		Action:    models.PaymentStatus(c.PostForm("action")),
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		appG.SendResponse(http.StatusBadRequest, common.ERROR_INVALID_PARAMS, nil, nil)
		return
	}

	err := models.ReviewStatus(payload.PaymentId, payload.Action)

	if err != nil {
		appG.SendResponse(http.StatusNotFound, common.ERROR_PAYMENT_NOT_FOUND, nil, err)
		return
	}

	res, _ := models.GetPaymentById(payload.PaymentId)

	appG.SendResponse(http.StatusAccepted, common.SUCCESS_OK, res, nil)
}
