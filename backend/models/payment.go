package models

import (
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/util"
	"github.com/akasection/durianpay-cs-dashboard/backend/services"
)

type Payment struct {
	PaymentID    uint          `gorm:"primaryKey" json:"payment_id"`
	MerchantName string        `json:"merchant_name"`
	Date         string        `json:"date"` // ISO 8601 format pls
	Amount       uint32        `json:"amount"`
	Status       PaymentStatus `json:"status"`
}

type PaymentStatus string

const (
	PaymentStatusCompleted  PaymentStatus = "completed"
	PaymentStatusProcessing PaymentStatus = "processing"
	PaymentStatusFailed     PaymentStatus = "failed"
)

type SortBy string
type OrderType string

const (
	SortByDate   SortBy = "date"
	SortByAmount SortBy = "amount"
)

const (
	OrderTypeAscending  OrderType = " ASC"
	OrderTypeDescending OrderType = " DESC"
)

func ListPayments(page int, limit int, status string, sortBy SortBy, orderType OrderType) ([]*Payment, error) {
	var payments []*Payment
	queryLimit := 10

	// if defined, set the query limit to custom value
	if limit > 0 {
		queryLimit = limit
	}

	queryLimit = util.ClampInt(queryLimit, 1, 100)
	offsetPage := util.ClampInt(page, 1, 32767)

	query := services.DB.Offset(offsetPage * queryLimit).Limit(queryLimit)

	if status != "" {
		query = query.Where("status = ?", status)
	}

	if sortBy == "" {
		sortBy = SortByDate
		orderType = OrderTypeDescending
	}
	query = query.Order(string(sortBy) + " " + string(orderType))

	err := query.Find(&payments).Error
	return payments, err
}

func GetPaymentById(id uint) (*Payment, error) {
	var payment Payment
	err := services.DB.First(&payment, id).Error
	if err != nil {
		return nil, err
	}
	return &payment, nil
}

func UpdatePaymentStatus(id uint, status string) error {
	return services.DB.Model(&Payment{}).Where("id = ?", id).Update("status", status).Error
}

func CountPaymentsByStatus(status string) (int64, error) {
	var count int64
	err := services.DB.Model(&Payment{}).Where("status = ?", status).Count(&count).Error
	return count, err
}
