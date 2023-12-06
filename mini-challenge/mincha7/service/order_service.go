package service

import (
	"mincha7/data/request"
	"mincha7/data/response"
)

type OrderService interface {
	Create(orders request.CreateOrderRequest)
	Update(orders request.UpdateOrderRequest) error
	Delete(ordersId int) error
	FindById(ordersId int) (response.OrderResponse, error)
	FindAll() []response.OrderResponse
}
