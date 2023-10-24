package service

import (
	"mincha7/repository"
)

type ItemServiceImpl struct {
	ItemRepository repository.ItemRepository
}

func NewItemServiceImpl(itemRepository repository.ItemRepository) ItemService {
	return &ItemServiceImpl{ItemRepository: itemRepository}
}

func (i ItemServiceImpl) DeleteByOrder(orderID int) {
	i.ItemRepository.DeleteByOrder(orderID)
}
