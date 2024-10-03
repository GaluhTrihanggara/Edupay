package model

import "time"

type Payment struct {
	UUIDPrimaryKey
	BillId      string    `gorm:"type:varchar(100)" json:"bill_id"`
	Amount      float64   `gorm:"type:numeric" json:"amoun"`
	Method      string    `gorm:"type:varchar(100)" json:"method"`
	Status      string    `gorm:"type:varchar(100)" json:"status"`
	PaymentDate time.Time `json:"payment_date"`
}
