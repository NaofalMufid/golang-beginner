package service

import (
	"mincha7/data/request"
	"mincha7/data/response"
)

type OrderService interface {
	Create(orders request.CreateOrderRequest)
	Update(orders request.UpdateOrderRequest)
	Delete(ordersId int)
	FindById(ordersId int) response.OrderResponse
	FindAll() []response.OrderResponse
}
