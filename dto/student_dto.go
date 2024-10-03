package dto

type StudentResponseDTO struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Class  string `json:"class"`
	Parent string `json:"parent"`
}
