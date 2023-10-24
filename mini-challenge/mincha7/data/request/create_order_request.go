package request

import "time"

type CreateOrderRequest struct {
	CustomerName string              `validate:"required" json:"customer_name"`
	OrderedAt    *time.Time          `validate:"required" json:"ordered_at"`
	Items        []CreateItemRequest `json:"items"`
}
