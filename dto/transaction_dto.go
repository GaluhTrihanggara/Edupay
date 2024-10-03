package dto

type TransactionRequestDTO struct {
	ItemId   string `json:"item_id" validate:"required"`
	Quantity int    `json:"quantity" validate:"required,min=1"`
}

type TransactionResponseDTO struct {
	Id              string  `json:"id"`
	ItemName        string  `json:"item_name"`
	Quantity        int     `json:"quantity"`
	TotalPrice      float64 `json:"total_price"`
	TransactionDate string  `json:"transaction_date"`
}
