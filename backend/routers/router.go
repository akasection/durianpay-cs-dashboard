package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/akasection/durianpay-cs-dashboard/backend/middleware/jwt"
	PaymentApi "github.com/akasection/durianpay-cs-dashboard/backend/routers/api/dashboard/v1"
	AuthApi "github.com/akasection/durianpay-cs-dashboard/backend/routers/api/dashboard/v1/auth"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger()) // Supply to logger vendor later
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	// Setup API v1
	apiv1 := r.Group("/dashboard/v1")
	apiv1.Use(jwt.JwtMiddleware())
	{
		// auth
		apiv1.POST("/auth/login", AuthApi.PostLogin)

		// dashboards
		apiv1.GET("/payments", PaymentApi.ListPayments)
		// apiv1.POST("/payments")
		apiv1.GET("/payment/:id/review")
		apiv1.POST("/payment/:id/review")
	}

	return r
}
