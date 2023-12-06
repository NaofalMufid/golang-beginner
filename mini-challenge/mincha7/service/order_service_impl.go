package service

import (
	"mincha7/data/request"
	"mincha7/data/response"
	"mincha7/helper"
	"mincha7/model"
	"mincha7/repository"

	"github.com/go-playground/validator/v10"
)

type OrderServiceImpl struct {
	OrdersRepository repository.OrdersRepository
	ItemService      ItemService
	Validate         *validator.Validate
}

func NewOrderServiceImpl(ordersRepository repository.OrdersRepository, itemService ItemService, validate *validator.Validate) OrderService {
	return &OrderServiceImpl{
		OrdersRepository: ordersRepository,
		ItemService:      itemService,
		Validate:         validate,
	}
}

func (o OrderServiceImpl) Create(order request.CreateOrderRequest) {
	err := o.Validate.Struct(order)
	helper.ErrorPanic(err)
	orderModel := model.Orders{
		CustomerName: order.CustomerName,
		OrderedAt:    order.OrderedAt,
	}

	var items []model.Items
	for _, itemReq := range order.Items {
		item := model.Items{
			Name:        itemReq.Name,
			Description: itemReq.Description,
			Quantity:    itemReq.Quantity,
		}
		item.OrderID = orderModel.Id
		items = append(items, item)
	}
	orderModel.Items = items
	o.OrdersRepository.Save(orderModel)
}

func (o OrderServiceImpl) Update(order request.UpdateOrderRequest) error {
	orderData, err := o.OrdersRepository.FindById(order.Id)
	if err != nil {
		return err
	}
	orderData.CustomerName = order.Customer_Name
	orderData.OrderedAt = order.Ordered_At
	o.OrdersRepository.Update(orderData)

	o.ItemService.DeleteByOrder(orderData.Id)
	for _, newItem := range order.Items {
		item := model.Items{
			Name:        newItem.Name,
			Description: newItem.Description,
			Quantity:    newItem.Quantity,
			OrderID:     orderData.Id,
		}
		o.ItemService.Create(item)
	}
	return nil
}

func (o OrderServiceImpl) Delete(orderId int) error {
	orderData, err := o.OrdersRepository.FindById(orderId)
	if err != nil {
		return err
	}
	// delete item
	o.ItemService.DeleteByOrder(orderData.Id)
	// delete order
	o.OrdersRepository.Delete(orderData.Id)
	return nil
}

func (o OrderServiceImpl) FindById(orderId int) (response.OrderResponse, error) {
	orderData, err := o.OrdersRepository.FindById(orderId)
	if err != nil {
		return response.OrderResponse{}, err
	}

	orderResponse := response.OrderResponse{
		Id:            orderData.Id,
		Customer_Name: orderData.CustomerName,
		Ordered_At:    orderData.OrderedAt,
	}
	for _, v := range orderData.Items {
		itemResponse := response.ItemResponse{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
			Quantity:    v.Quantity,
		}
		orderResponse.Items = append(orderResponse.Items, itemResponse)
	}
	return orderResponse, nil
}

func (o OrderServiceImpl) FindAll() []response.OrderResponse {
	result := o.OrdersRepository.FindAll()

	var orders []response.OrderResponse
	for _, v := range result {
		order := response.OrderResponse{
			Id:            v.Id,
			Customer_Name: v.CustomerName,
			Ordered_At:    v.OrderedAt,
		}
		for _, item := range v.Items {
			itemResponse := response.ItemResponse{
				Id:          item.Id,
				Name:        item.Name,
				Description: item.Description,
				Quantity:    item.Quantity,
			}
			order.Items = append(order.Items, itemResponse)
		}
		orders = append(orders, order)
	}
	return orders
}
