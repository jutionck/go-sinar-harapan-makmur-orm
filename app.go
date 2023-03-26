package main

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/config"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	dbConn, _ := config.NewDbConnection(c)
	db := dbConn.Conn()

	// dbConn.Migrate(&model.Customer{}, &model.UserCredential{})
	// bod, _ := time.Parse("2006-01-11", "1995-03-07")
	// customer := &model.Customer{
	// 	FirstName:   "Jution",
	// 	LastName:    "Candra",
	// 	Address:     "Jakarta Selatan",
	// 	Email:       "jutionck@gmail.com",
	// 	PhoneNumber: "028292992",
	// 	Bod:         bod,
	// 	UserCredential: model.UserCredential{
	// 		UserName: "jutionck@gmail.com",
	// 		Password: "password",
	// 		IsActive: false,
	// 	},
	// }

	// // if err := db.Create(customer); err != nil {
	// // 	fmt.Println(err)
	// // }

	// // Update customer address
	// customer.ID = "58427210-9988-47c1-8859-67e2bc9b0ee4"
	// customer.Address = "Pasar Minggu"
	// if err := db.Updates(&customer); err != nil {
	// 	fmt.Println(err)
	// }

	// get all customer
	// var customers []model.Customer
	// if err := db.Find(&customers); err != nil {
	// 	fmt.Println(err)
	// }

	// for _, c := range customers {
	// 	fmt.Println("customer:", c)
	// 	fmt.Println("user credential:", c.UserCredential)
	// 	fmt.Println()
	// }

	// Joins
	// var customers []model.Customer
	// if err := db.Joins("UserCredential").Find(&customers).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// for _, c := range customers {
	// 	fmt.Println("customer:", c.FirstName, c.LastName, c.UserCredentialID)
	// 	fmt.Println("user credential:", c.UserCredential.UserName)
	// 	fmt.Println()
	// }

	// Select
	// var customers []model.Customer
	// if err := db.Model(&model.Customer{}).Select("first_name, last_name, user_credential_id").Joins("UserCredential").Scan(&customers).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// for _, c := range customers {
	// 	fmt.Println("customer:", c.FirstName, c.LastName, c.UserCredentialID)
	// 	fmt.Println("user name:", c.UserCredential.UserName)
	// 	fmt.Println()
	// }

	// get customer + user credential (LAZY)
	// var customer model.Customer
	// if err := db.First(&customer, "id = ?", "58427210-9988-47c1-8859-67e2bc9b0ee4").Error; err != nil {
	// 	panic(err)
	// }

	// // Load user credential data only when accessed
	// if err := db.Model(&customer).Association("UserCredential").Find(&customer.UserCredential); err != nil {
	// 	panic(err)
	// }

	// fmt.Println("customer:", customer.FirstName, customer.LastName, customer.UserCredentialID)
	// fmt.Println("user credential:", customer.UserCredential)
	// fmt.Println()

	// Preload
	// var customers []model.Customer
	// if err := db.Preload("UserCredential").Find(&customers).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// for _, c := range customers {
	// 	fmt.Println("customer:", c.FirstName, c.LastName, c.UserCredentialID)
	// 	fmt.Println("user name:", c.UserCredential.UserName)
	// 	fmt.Println()
	// }

	// Get Customer
	// var customer01 model.Customer
	// if err := db.First(&customer01, "id=?", "58427210-9988-47c1-8859-67e2bc9b0ee4"); err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(customer01.UserCredentialID)

	// Update user credential -> Password
	// uc := model.UserCredential{Password: "inipasswordkuat"}
	// customer01.UserCredential = uc
	// if err := db.Model(&model.UserCredential{}).Where("id=?", customer01.UserCredentialID).Updates(&customer01.UserCredential).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// Get Customer + User Credential
	var customer01 model.Customer
	if err := db.Preload("UserCredential").First(&customer01, "id=?", "58427210-9988-47c1-8859-67e2bc9b0ee4"); err != nil {
		fmt.Println(err)
	}

	fmt.Println(customer01.UserCredential.ID)

	// // Update user credential -> Password
	uc := model.UserCredential{Password: "12345"}
	if err := db.Model(&customer01.UserCredential).Updates(&uc).Error; err != nil {
		fmt.Println(err)
	}
}
