package routers

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// "github.com/akasection/durianpay-cs-dashboard/backend/middleware/jwt"
	"github.com/akasection/durianpay-cs-dashboard/backend/middleware"
	"github.com/akasection/durianpay-cs-dashboard/backend/pkg/dto"
	PaymentApi "github.com/akasection/durianpay-cs-dashboard/backend/routers/api/dashboard/v1"
	AuthApi "github.com/akasection/durianpay-cs-dashboard/backend/routers/api/dashboard/v1/auth"
	PaymentReviewApi "github.com/akasection/durianpay-cs-dashboard/backend/routers/api/dashboard/v1/payment/_id"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// r.Use(gin.Logger()) // Supply to logger vendor later
	// r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{os.Getenv("FRONTEND_URL")},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Setup login endpoint
	r.POST("/dashboard/v1/auth/login", AuthApi.PostLogin)

	// Setup API v1
	apiv1 := r.Group("/dashboard/v1")
	apiv1.Use(middleware.UseJwt())
	{
		// auth

		// dashboards
		apiv1.GET("/payments", middleware.UseRbac(dto.ROLE_CS, dto.ROLE_OPERATION), PaymentApi.GetListPayments)
		apiv1.PUT("/payment/:id/review", middleware.UseRbac(dto.ROLE_OPERATION), PaymentReviewApi.PutReviewPayment)
	}

	return r
}
