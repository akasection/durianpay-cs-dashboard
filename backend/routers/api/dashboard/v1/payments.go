package PaymentApi

import (
	"strconv"

	PaymentModel "github.com/akasection/durianpay-cs-dashboard/backend/models"
	"github.com/gin-gonic/gin"
)

func ListPayments(c *gin.Context) {

	// Validate query parameters
	status := c.Query("status")
	limit, _ := strconv.Atoi(c.Query("limit"))
	sort := c.Query("sort")
	order := c.Query("order")
	page, _ := strconv.Atoi(c.Query("page"))

	payments, pErr := PaymentModel.ListPayments(page, limit, status, PaymentModel.SortBy(sort), PaymentModel.OrderType(order))

	if pErr != nil {
		c.JSON(500, gin.H{"error": pErr.Error()})
		return
	}
	c.JSON(200, payments)
}
