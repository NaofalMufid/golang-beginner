package response

import "time"

type OrderResponse struct {
	Id            int            `json:"id"`
	Customer_Name string         `json:"customer_name"`
	Ordered_At    *time.Time     `json:"ordered_at"`
	Items         []ItemResponse `json:"items"`
}
