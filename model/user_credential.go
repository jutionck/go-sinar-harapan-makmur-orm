package model

type UserCredential struct {
	BaseModel
	UserName   string `gorm:"unique;size:50;not null"`
	Password   string `gorm:"size:60;not null"`
	IsActive   bool   `gorm:"default:false"`
	CustomerID string
}

func (UserCredential) TableName() string {
	return "mst_user"
}
