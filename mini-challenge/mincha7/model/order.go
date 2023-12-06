package model

import (
	"time"
)

type Orders struct {
	Id           int        `gorm:"type:int;primaryKey;autoIncrement"`
	CustomerName string     `gorm:"type:varchar(100)"`
	OrderedAt    *time.Time `gorm:"type:timestamp"`
	CreatedAt    *time.Time `gorm:"type:timestamp"`
	UpdatedAt    *time.Time `gorm:"type:timestamp"`
	Items        []Items    `gorm:"ForeignKey:OrderID;AssociationForeignKey:ID"`
}
