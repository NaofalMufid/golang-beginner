package repository

import (
	"mincha7/helper"
	"mincha7/model"

	"gorm.io/gorm"
)

type ItemRepositoryImpl struct {
	Db *gorm.DB
}

func NewItemRepository(Db *gorm.DB) ItemRepository {
	return &ItemRepositoryImpl{Db: Db}
}

func (i ItemRepositoryImpl) Save(item model.Items) error {
	result := i.Db.Create(&item)
	return result.Error
}

func (i ItemRepositoryImpl) Update(item model.Items) error {
	result := i.Db.Save(&item)
	return result.Error
}

func (i ItemRepositoryImpl) DeleteByOrder(orderID int) {
	var item model.Items
	result := i.Db.Where("order_id = ?", orderID).Delete(&item)
	helper.ErrorPanic(result.Error)
}

func (i ItemRepositoryImpl) FindByOrderAndItemID(orderID, itemID int) (model.Items, error) {
	var item model.Items
	result := i.Db.Where("order_id = ? AND id = ?", orderID, itemID).First(&item)
	if result.Error != nil {
		return item, result.Error
	}
	return item, nil
}
