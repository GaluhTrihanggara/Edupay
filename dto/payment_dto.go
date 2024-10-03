package dto

type PaymentRequestDTO struct {
	BillId string  `json:"bill_id" validate:"required"`
	Amount float64 `json:"amount" validate:"required,numeric"`
	Method string  `json:"method" validate:"required"`
}

type PaymentResponseDTO struct {
	Id          string  `json:"id"`
	BillId      string  `json:"bill_id"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
	Method      string  `json:"method"`
	PaymentDate string  `json:"payment_date"`
}
