package PaymentApi

import (
	"net/http"
	"strconv"

	PaymentModel "github.com/akasection/durianpay-cs-dashboard/backend/models"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/common"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/ginutil"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/util"
	"github.com/gin-gonic/gin"
)

type ListPaymentsQuery struct {
	Sort   PaymentModel.SortBy    `form:"sort" binding:"omitempty,oneof=date amount status"`
	Order  PaymentModel.OrderType `form:"order" binding:"omitempty,oneof=asc desc"`
	Page   int                    `form:"page" binding:"omitempty,min=1"`
	Status string                 `form:"status" binding:"omitempty,oneof=processing completed failed"`
	Limit  int                    `form:"limit" binding:"omitempty,min=1,max=100"`
}

func GetListPayments(c *gin.Context) {
	appG := ginutil.Gin{C: c}

	// Validate query parameters
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))
	status := c.Query("status")
	sort := c.Query("sort")
	order := c.Query("order")
	var query = ListPaymentsQuery{
		Sort:   PaymentModel.SortBy(sort),
		Order:  PaymentModel.OrderType(order),
		Page:   page,
		Status: status,
		Limit:  limit,
	}
	if err := c.ShouldBind(&query); err != nil {
		// TODO: filter err based on validation error type
		appG.SendResponse(http.StatusBadRequest, common.ERROR_INVALID_PARAMS, nil, nil)
		return
	}

	payments, pErr := PaymentModel.ListPayments(page, limit, status, PaymentModel.SortBy(sort), PaymentModel.OrderType(order))
	counts := PaymentModel.CountTotalPayments()
	if pErr != nil {
		appG.SendResponse(http.StatusInternalServerError, common.ERROR_GENERIC, nil, nil)
		return
	}

	appG.SendResponse(http.StatusOK, common.SUCCESS_OK, payments, ginutil.Meta{
		Offset: util.ClampInt((page-1), 0, 32767) * len(payments),
		Limit:  len(payments),
		Total:  counts.Total,
		Extra: map[string]interface{}{
			"completed":  counts.Completed,
			"processing": counts.Processing,
			"failed":     counts.Failed,
		},
	})
}
