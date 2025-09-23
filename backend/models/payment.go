package models

import (
	"context"
	"time"

	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/util"
	"github.com/akasection/durianpay-cs-dashboard/backend/services"
	"gorm.io/gorm"
)

type Payment struct {
	PaymentID    uint          `gorm:"primaryKey" json:"payment_id"`
	MerchantName string        `json:"merchant_name"`
	Date         time.Time     `gorm:"type:datetime;default:CURRENT_TIMESTAMP" json:"date"` // ISO 8601 format pls
	Amount       uint32        `json:"amount"`
	Status       PaymentStatus `json:"status"`
}

type PaymentStatus string

const (
	PaymentStatusCompleted  PaymentStatus = "completed"
	PaymentStatusProcessing PaymentStatus = "processing"
	PaymentStatusFailed     PaymentStatus = "failed"
)

type StatusCountMeta struct {
	Total      int `json:"total"`
	Completed  int `json:"completed"`
	Processing int `json:"processing"`
	Failed     int `json:"failed"`
}

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

	// limit will set to 10 or defined "limit"
	if limit > 0 {
		queryLimit = limit
	}

	queryLimit = util.ClampInt(queryLimit, 1, 100)
	offsetPage := util.ClampInt(page, 1, 32767)

	query := services.DB.Offset((offsetPage - 1) * queryLimit).Limit(queryLimit)

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

func ReviewStatus(id uint, status PaymentStatus) error {
	ctx := context.Background()

	if status == PaymentStatusProcessing {
		return gorm.ErrInvalidData
	}

	result, err := gorm.G[Payment](services.DB).
		Where("payment_id = ? AND status = ?", id, PaymentStatusProcessing).
		Update(ctx, "status", status)

	if err != nil {
		return err
	}
	if result == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}

func CountTotalPayments() StatusCountMeta {
	ctx := context.Background()
	base := gorm.G[Payment](services.DB)
	totalCount, _ := base.Count(ctx, "payment_id")
	completedCount, _ := base.Where("status = ?", PaymentStatusCompleted).Count(ctx, "payment_id")
	processingCount, _ := base.Where("status = ?", PaymentStatusProcessing).Count(ctx, "payment_id")
	failedCount, _ := base.Where("status = ?", PaymentStatusFailed).Count(ctx, "payment_id")

	return StatusCountMeta{
		Total:      int(totalCount),
		Completed:  int(completedCount),
		Processing: int(processingCount),
		Failed:     int(failedCount),
	}
}
