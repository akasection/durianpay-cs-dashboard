package PaymentReviewApi

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/akasection/durianpay-cs-dashboard/backend/models"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/common"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/ginutil"
)

type ReviewAction struct {
	Action models.PaymentStatus `json:"action" binding:"required,oneof=completed failed"`
}

func PutReviewPayment(c *gin.Context) {
	appG := ginutil.Gin{C: c}
	paymentId, idErr := strconv.Atoi(c.Param("id"))

	if idErr != nil || paymentId <= 0 {
		appG.SendResponse(http.StatusBadRequest, common.ERROR_INVALID_PARAMS, nil, nil)
		return
	}

	var payload ReviewAction
	if bindErr := c.ShouldBindJSON(&payload); bindErr != nil {
		appG.SendResponse(http.StatusBadRequest, common.ERROR_INVALID_PARAMS, nil, nil)
		return
	}

	err := models.ReviewStatus(uint(paymentId), payload.Action)

	if err != nil {
		appG.SendResponse(http.StatusNotFound, common.ERROR_PAYMENT_NOT_FOUND, nil, err)
		return
	}

	res, _ := models.GetPaymentById(uint(paymentId))

	appG.SendResponse(http.StatusOK, common.SUCCESS_OK, res, nil)
}
