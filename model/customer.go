package model

import "time"

type Customer struct {
	BaseModel
	FirstName      string `gorm:"size:30"`
	LastName       string `gorm:"size:30"`
	Address        string
	Email          string `gorm:"unique;size:30"`
	PhoneNumber    string `gorm:"unique;size:15"`
	Bod            time.Time
	UserCredential UserCredential
}

func (Customer) TableName() string {
	return "mst_customer"
}
