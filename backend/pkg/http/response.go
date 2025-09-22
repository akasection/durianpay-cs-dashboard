package http

import (
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/common"
	"github.com/gin-gonic/gin"
)

// Standard API response structure
type ApiResponse struct {
	Code    int         `json:"code"`    // e.g., 200 for success, 400 for client error, etc.
	Message string      `json:"message"` // A human-readable message
	Data    interface{} `json:"data"`    // The actual data payload
}

type Gin struct {
	C *gin.Context
}

func (g *Gin) SendResponse(httpCode, errCode int, data interface{}) {
	g.C.JSON(httpCode, ApiResponse{
		Code:    errCode,
		Message: common.MessageCode(errCode),
		Data:    data,
	})
}
