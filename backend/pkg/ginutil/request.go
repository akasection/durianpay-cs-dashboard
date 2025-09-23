package ginutil

// Common request query structure

type PaginationQuery struct {
	Page     int `form:"page" json:"page" binding:"required,min=1"`
	PageSize int `form:"page_size" json:"page_size" binding:"required,min=1,max=100"`
}
