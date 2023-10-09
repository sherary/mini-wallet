package routes

import (
	controller "mini-wallet/controllers"

	"github.com/gin-gonic/gin"
)

func IndexRouter(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.POST("/wallet/init", controller.Signup)
}
