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
