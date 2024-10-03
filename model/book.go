package model

type Book struct {
	UUIDPrimaryKey
	Title  string  `gorm:"type:varchar(100)" json:"title"`
	Author string  `gorm:"type:varchar(100)" json:"author"`
	Price  float64 `gorm:"type:numeric" json:"price"`
	Stock  int     `gorm:"type:int" json:"stock"`
}
