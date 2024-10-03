package model

type Student struct {
	UUIDPrimaryKey
	Name     string `gorm:"type:varchar(100)" json:"name"`
	Class    string `gorm:"type:varchar(50)" json:"class"`
	ParentId string `gorm:"type:varchar(50)" json:"parent_id"`
	Parent   User   `gorm:"foreigntKey:ParentID" json:"parent"`
}
