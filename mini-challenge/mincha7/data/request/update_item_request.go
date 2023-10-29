package request

type UpdateItemRequest struct {
	Name        string `validate:"required" json:"name"`
	Description string `validate:"required" json:"description"`
	Quantity    int    `validate:"required" json:"quantity"`
}
