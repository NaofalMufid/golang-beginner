package response

type ItemResponse struct {
	Id          int    `json:"id"`
	Name        string `json:"customer_name"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}
