package model

type Shirt struct {
	UUIDPrimaryKey
	Name  string  `gorm:"type:varchar(100)" json:"name"`
	Size  string  `gorm:"type:varchar(100)" json:"size"`
	Price float64 `gorm:"type:varchar(100)" json:"price"`
	Stock int     `gorm:"type:int" json:"stock"`
}
