package model

import "time"

type PaymentHistory struct {
	UUIDPrimaryKey
	ParentId    string    `json:"parent_id"`
	PaymentId   string    `json:"payment_id"`
	Amount      float64   `gorm:"type:numeric" json:"amount"`
	Status      string    `gorm:"type:varchar(100)" json:"status"`
	PaymentDate time.Time `json:"payment_date"`
}
