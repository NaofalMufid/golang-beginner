package request

import "time"

type UpdateOrderRequest struct {
	Id            int                 `validate:"required"`
	Customer_Name string              `validate:"required" json:"customer_name"`
	Ordered_At    *time.Time          `validate:"required" json:"ordered_at"`
	Items         []UpdateItemRequest `json:"items"`
}

type UpdateOrderData struct {
	Id            int        `validate:"required"`
	Customer_Name string     `validate:"required" json:"customer_name"`
	Ordered_At    *time.Time `validate:"required" json:"ordered_at"`
}
