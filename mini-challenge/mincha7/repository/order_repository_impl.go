package repository

import (
	"errors"
	"mincha7/data/request"
	"mincha7/helper"
	"mincha7/model"

	"gorm.io/gorm"
)

type OrdersRepositoryImpl struct {
	Db *gorm.DB
}

func NewOrdersRepository(Db *gorm.DB) OrdersRepository {
	return &OrdersRepositoryImpl{Db: Db}
}

func (o OrdersRepositoryImpl) Save(orders model.Orders) {
	result := o.Db.Create(&orders)
	helper.ErrorPanic(result.Error)
}

func (o OrdersRepositoryImpl) Update(orders model.Orders) {
	var updateOrder = request.UpdateOrderRequest{
		Id:            orders.Id,
		Customer_Name: orders.CustomerName,
		Ordered_At:    orders.OrderedAt,
	}

	result := o.Db.Model(&orders).Updates(updateOrder)
	helper.ErrorPanic(result.Error)
}

func (o OrdersRepositoryImpl) Delete(ordersId int) {
	var orders model.Orders
	result := o.Db.Where("id = ?", ordersId).Delete(&orders)
	helper.ErrorPanic(result.Error)
}

func (o OrdersRepositoryImpl) FindById(ordersId int) (model.Orders, error) {
	var order model.Orders
	result := o.Db.Preload("Items").Find(&order, ordersId)
	if result != nil {
		return order, nil
	} else {
		return order, errors.New("order not found")
	}
}

func (o OrdersRepositoryImpl) FindAll() []model.Orders {
	var orders []model.Orders
	result := o.Db.Preload("Items").Find(&orders)
	helper.ErrorPanic(result.Error)
	return orders
}
