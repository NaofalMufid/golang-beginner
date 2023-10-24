package request

type CreateItemRequest struct {
	Name        string `validate:"required" json:"name"`
	Description string `validate:"required" json:"description"`
	Quantity    int    `validate:"required" json:"quantity"`
}
