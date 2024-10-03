package dto

type PaginationRequestDTO struct {
	Page    int `json:"page" validate:"required,min=1"`
	PerPage int `json:"per_page" validate:"required,min=1"`
}

type PaginationResponseDTO struct {
	Page       int         `json:"page"`
	PerPage    int         `json:"per_page"`
	TotalPages int         `json:"total_pages"`
	TotalItems int         `json:"total_items"`
	Data       interface{} `json:"data"`
}
