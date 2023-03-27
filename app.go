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

	// if err := db.Create(customer); err != nil {
	// 	fmt.Println(err)
	// }

	// update the user credential for a customer
	// customer := &model.Customer{}
	// if err := db.Preload("UserCredential").First(customer, "id = ?", "b135e13d-abaa-4c49-a246-7957b5ed81d6").Error; err != nil {
	// 	panic(err)
	// }
	// fmt.Println(customer.UserCredential)

	// customer.UserCredential.UserName = "jutionck@gmail.com"
	// customer.UserCredential.Password = "newpassword"
	// if err := db.Save(customer.UserCredential).Error; err != nil {
	// 	panic(err)
	// }

	// delete a customer and their associated user credential
	// customer := &model.Customer{}
	// if err := db.Preload("UserCredential").First(customer, "id = ?", "b135e13d-abaa-4c49-a246-7957b5ed81d6").Error; err != nil {
	// 	panic(err)
	// }
	// tx := db.Begin()
	// if err := tx.Delete(customer).Error; err != nil {
	// 	panic(err)
	// }

	// if err := tx.Delete(&customer.UserCredential).Error; err != nil {
	// 	panic(err)
	// }
	// tx.Commit()

	// retrieve a customer with their associated user credential
	customer := &model.Customer{}
	if err := db.Preload("UserCredential").First(customer, "id = ?", "b135e13d-abaa-4c49-a246-7957b5ed81d6").Error; err != nil {
		panic(err)
	}
	fmt.Println("Customer:", customer)
	fmt.Println("UserCredential:", customer.UserCredential)

}
