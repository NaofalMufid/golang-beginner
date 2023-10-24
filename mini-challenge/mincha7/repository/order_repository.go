package repository

import "mincha7/model"

type OrdersRepository interface {
	Save(orders model.Orders)
	Update(orders model.Orders)
	Delete(ordersId int)
	FindById(ordersId int) (orders model.Orders, err error)
	FindAll() []model.Orders
}
