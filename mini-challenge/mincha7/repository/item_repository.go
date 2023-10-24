package repository

import "mincha7/model"

type ItemRepository interface {
	AddToOrder(orderID int, item model.Items)
	Update(item model.Items)
	DeleteByOrder(orderID int)
	FindByOrderAndItemID(orderID, itemID int) (model.Items, error)
}
