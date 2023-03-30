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
	// db := dbConn.Conn()

	dbConn.Migrate(&model.Customer{}, &model.Vehicle{})

	// Insert 1 By 1
	// password, err := utils.HashPassword("P@ssword")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// newCustomer := model.Customer{
	// 	FirstName:   "Jution",
	// 	LastName:    "Kirana",
	// 	Address:     "Jakarta Selatan",
	// 	Email:       "jution.kirana@enigmacamp.com",
	// 	PhoneNumber: "08219292929",
	// 	Bod:         time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC),
	// 	UserCredential: model.UserCredential{
	// 		UserName: "jution.kirana@enigmacamp.com",
	// 		Password: password,
	// 		IsActive: false,
	// 	},
	// }

	// vehicle := []model.Vehicle{
	// 	{
	// 		BrandID:        "1d6b839a-4678-4d36-9987-99c4d31449d8",
	// 		Model:          "Jazz",
	// 		ProductionYear: 2022,
	// 		Color:          "Kuning",
	// 		IsAutomatic:    true,
	// 		Stock:          10,
	// 		SalePrice:      300000000,
	// 		Status:         "baru",
	// 	},
	// 	{
	// 		BrandID:        "1d6b839a-4678-4d36-9987-99c4d31449d8",
	// 		Model:          "City",
	// 		ProductionYear: 2023,
	// 		Color:          "Hitam",
	// 		IsAutomatic:    true,
	// 		Stock:          3,
	// 		SalePrice:      400000000,
	// 		Status:         "baru",
	// 	},
	// }

	// db.Create(&newCustomer)
	// db.Create(&vehicle)

	// Insert to customer_vehicles
	// var customer model.Customer
	// if err := db.Where("id=?", "cacc52cb-7581-4787-b8b1-48b29a4b7219").First(&customer).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// // karena menggunakan pointer
	// var vehicles []*model.Vehicle
	// if err := db.Find(&vehicles).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// customer.Vehicles = vehicles
	// if err := db.Model(&customer).Updates(customer).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// // Create Customer + Vehicle
	// password, err := utils.HashPassword("P@ssword")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// newCustomer := model.Customer{
	// 	FirstName:   "Jution",
	// 	LastName:    "Kirana",
	// 	Address:     "Jakarta Selatan",
	// 	Email:       "jution.kirana@enigmacamp.com",
	// 	PhoneNumber: "08219292929",
	// 	Bod:         time.Date(1995, 1, 1, 0, 0, 0, 0, time.UTC),
	// 	UserCredential: model.UserCredential{
	// 		UserName: "jution.kirana@enigmacamp.com",
	// 		Password: password,
	// 		IsActive: false,
	// 	},
	// 	Vehicles: []*model.Vehicle{
	// 		{
	// 			BrandID:        "1d6b839a-4678-4d36-9987-99c4d31449d8",
	// 			Model:          "Jazz",
	// 			ProductionYear: 2022,
	// 			Color:          "Kuning",
	// 			IsAutomatic:    true,
	// 			Stock:          10,
	// 			SalePrice:      300000000,
	// 			Status:         "baru",
	// 		},
	// 		{
	// 			BrandID:        "1d6b839a-4678-4d36-9987-99c4d31449d8",
	// 			Model:          "City",
	// 			ProductionYear: 2023,
	// 			Color:          "Hitam",
	// 			IsAutomatic:    true,
	// 			Stock:          3,
	// 			SalePrice:      400000000,
	// 			Status:         "baru",
	// 		},
	// 	},
	// }

	// db.Create(&newCustomer)

	// Setelah itu kita akan menambahkan data di tabel pivot nya
	// Kita bisa memanfaatkan Association Mode

	// Buat Assocuation Mode
	// var customer model.Customer
	// var vehicle model.Vehicle

	// // Untuk menambahkan data ke dalam tabel pivot kita bisa menggunakan `Append`
	// // Pertama kita cari customer yang mau ditambah
	// if err := db.Where("id=?", "e066762a-16fa-4210-bded-60a34b449aa1").First(&customer).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// // Kedua kita cari vehicle yang mau ditambah
	// if err := db.Where("id=?", "ae4c2a63-ab2b-4692-affa-9061276c394a").First(&vehicle).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// // // Kemudian kita lakukan appen
	// if err := db.Model(&customer).Association("Vehicles").Append(&vehicle); err != nil {
	// 	fmt.Println(err)
	// }
}
