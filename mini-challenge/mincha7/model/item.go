package model

import "time"

type Items struct {
	Id          int        `gorm:"type:int;primaryKey;autoIncrement"`
	Name        string     `gorm:"type:varchar(255)"`
	Description string     `gorm:"type:text"`
	Quantity    int        `gorm:"type:int"`
	OrderID     int        `gorm:"type:int;ForeignKey:OrderID"`
	CreatedAt   *time.Time `gorm:"type:timestamp"`
	UpdatedAt   *time.Time `gorm:"type:timestamp"`
}
