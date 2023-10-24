package router

import (
	"mincha7/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter(orderController *controller.OrderController) *gin.Engine {
	service := gin.Default()

	service.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "Assigment REST API")
	})

	service.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page Not Found"})
	})

	router := service.Group("/api")
	orderRouter := router.Group("/order")
	orderRouter.GET("", orderController.FindAll)
	orderRouter.GET("/:order_id", orderController.FindById)
	orderRouter.POST("", orderController.Create)
	orderRouter.PUT("/:order_id", orderController.Update)
	orderRouter.DELETE("/:order_id", orderController.Delete)

	return service
}
