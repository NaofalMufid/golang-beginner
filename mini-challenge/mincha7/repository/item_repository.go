package repository

import "mincha7/model"

type ItemRepository interface {
	Save(item model.Items) error
	Update(item model.Items) error
	DeleteByOrder(orderID int)
	FindByOrderAndItemID(orderID, itemID int) (model.Items, error)
}
