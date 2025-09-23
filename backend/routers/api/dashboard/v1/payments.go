package PaymentApi

import (
	"net/http"

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

// @Summary Get list of payments
// @Produce  json
// @Param id path int true "ID"
// @Param sort query string false "Sort by field" Enums(date, amount, status) default(date)
// @Param order query string false "Order type" Enums(asc, desc) default(desc)
// @Param page query int false "Page number" minimum(1) default(1)
// @Param status query string false "Filter by status" Enums(processing, completed, failed)
// @Param limit query int false "Limit number of items per page" minimum(1) maximum(100) default(10)
// @Success 200 {string}  json "{ "data": [{"payment_id": 1, "merchant_name": "Merchant A", "date": "2023-10-01T12:00:00Z", "amount": 1000, "status": "completed"}] }"
// @Router /dashboard/v1/payments [get]
func GetListPayments(c *gin.Context) {
	appG := ginutil.Gin{C: c}
	var query ListPaymentsQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		// TODO: filter err based on validation error type
		appG.SendResponse(http.StatusBadRequest, common.ERROR_INVALID_PARAMS, nil, nil)
		return
	}

	payments, pErr := PaymentModel.ListPayments(query.Page, query.Limit, query.Status, PaymentModel.SortBy(query.Sort), PaymentModel.OrderType(query.Order))
	counts := PaymentModel.CountTotalPayments()
	if pErr != nil {
		appG.SendResponse(http.StatusInternalServerError, common.ERROR_GENERIC, nil, nil)
		return
	}

	var countIndex = map[string]interface{}{
		"all":        counts.Total,
		"completed":  counts.Completed,
		"processing": counts.Processing,
		"failed":     counts.Failed,
	}

	countLabel := "all"
	if query.Status != "" {
		countLabel = query.Status
	}

	appG.SendResponse(http.StatusOK, common.SUCCESS_OK, payments, ginutil.Meta{
		Offset: util.ClampInt((query.Page-1), 0, 32767) * query.Limit,
		Limit:  query.Limit,
		Total:  countIndex[countLabel].(int),
		Extra:  countIndex,
	})
}
