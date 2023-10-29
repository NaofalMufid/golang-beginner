package service

import (
	"mincha7/model"
)

type ItemService interface {
	Create(item model.Items)
	Update(item model.Items) error
	DeleteByOrder(orderID int)
	FindByOrderAndItemID(orderID, itemID int) (model.Items, error)
}
