package controller

import (
	"mincha7/data/request"
	"mincha7/data/response"
	"mincha7/helper"
	"mincha7/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService service.OrderService
}

func NewOrderController(service service.OrderService) *OrderController {
	return &OrderController{orderService: service}
}

func (controller *OrderController) Create(ctx *gin.Context) {
	createOrderRequest := request.CreateOrderRequest{}
	err := ctx.ShouldBindJSON(&createOrderRequest)
	helper.ErrorPanic(err)

	controller.orderService.Create(createOrderRequest)

	webResponse := response.Response{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *OrderController) Update(ctx *gin.Context) {
	updateOrderRequest := request.UpdateOrderRequest{}
	err := ctx.ShouldBindJSON(&updateOrderRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orderId := ctx.Param("order_id")
	id, err := strconv.Atoi(orderId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	updateOrderRequest.Id = id

	err = controller.orderService.Update(updateOrderRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *OrderController) Delete(ctx *gin.Context) {
	orderId := ctx.Param("order_id")
	id, err := strconv.Atoi(orderId)
	helper.ErrorPanic(err)

	err = controller.orderService.Delete(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:   200,
		Status: "OK",
		Data:   nil,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *OrderController) FindById(ctx *gin.Context) {
	orderId := ctx.Param("order_id")
	id, err := strconv.Atoi(orderId)
	helper.ErrorPanic(err)

	orderResponse, err := controller.orderService.FindById(id)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := response.Response{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *OrderController) FindAll(ctx *gin.Context) {
	orderResponse := controller.orderService.FindAll()

	webResponse := response.Response{
		Code:   200,
		Status: "OK",
		Data:   orderResponse,
	}
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, webResponse)
}
