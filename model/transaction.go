package model

import "time"

type Transaction struct {
	UUIDPrimaryKey
	ParentId        string    `gorm:"type:varchar(100)" json:"parent_id"`
	ItemId          string    `gorm:"type:varchar(100)" json:"item_id"`
	Quantity        string    `gorm:"type:varchar(100)" json:"quantity"`
	TotalPrice      float64   `gorm:"type:numeric" json:"total_price"`
	TransactionDate time.Time `json:"transaction_date"`
}
