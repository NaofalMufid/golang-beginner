package service

import (
	"mincha7/helper"
	"mincha7/model"
	"mincha7/repository"
)

type ItemServiceImpl struct {
	ItemRepository repository.ItemRepository
}

func NewItemServiceImpl(itemRepository repository.ItemRepository) ItemService {
	return &ItemServiceImpl{ItemRepository: itemRepository}
}

func (i ItemServiceImpl) Create(item model.Items) {
	itemModel := model.Items{
		Name:        item.Name,
		Description: item.Description,
		Quantity:    item.Quantity,
		OrderID:     item.OrderID,
	}
	err := i.ItemRepository.Save(itemModel)
	helper.ErrorPanic(err)
}

func (i ItemServiceImpl) Update(item model.Items) error {
	err := i.ItemRepository.Update(item)
	if err != nil {
		return err
	}
	return nil
}

func (i ItemServiceImpl) DeleteByOrder(orderID int) {
	i.ItemRepository.DeleteByOrder(orderID)
}

func (i ItemServiceImpl) FindByOrderAndItemID(orderID, itemID int) (model.Items, error) {
	item, err := i.ItemRepository.FindByOrderAndItemID(orderID, itemID)
	if err != nil {
		return model.Items{}, err
	}
	return item, nil
}
