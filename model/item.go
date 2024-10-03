package model

type Item struct {
	UUIDPrimaryKey
	Name        string  `gorm:"type:varchar(100)" json:"name"`
	Description string  `gorm:"type:varchar(100)" json:"description"`
	Price       float64 `gorm:"type:numeric" json:"price"`
	Stock       string  `gorm:"type:varchar(100)" json:"stock"`
}
