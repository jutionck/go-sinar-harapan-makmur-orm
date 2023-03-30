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

	dbConn.Migrate(
		&model.UserCredential{},
		&model.Customer{},
		&model.Brand{},
		&model.Vehicle{},
		&model.Employee{},
		&model.Transaction{},
	)

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

	// newBrand := model.Brand{
	// 	Name: "Honda",
	// 	Vehicle: []model.Vehicle{
	// 		{
	// 			Model:          "Jazz",
	// 			ProductionYear: 2022,
	// 			Color:          "Kuning",
	// 			IsAutomatic:    true,
	// 			Stock:          10,
	// 			SalePrice:      300000000,
	// 			Status:         "baru",
	// 		},
	// 		{
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

	// db.Debug().Create(&newCustomer)
	// db.Debug().Create(&newBrand)

	// Insert to customer_vehicles
	// var vehicle model.Vehicle
	// if err := db.Where("id=?", "87d8ebc0-b18b-44ae-9e40-7f8e7198a53f").First(&vehicle).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// var customers []model.Customer
	// if err := db.Find(&customers).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// vehicle.Customers = customers
	// if err := db.Debug().Model(&vehicle).Updates(&vehicle).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// // Create Customer + Vehicle
	// password, err := utils.HashPassword("P@ssword")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// newCustomerVehicle := model.Customer{
	// 	FirstName:   "Edo",
	// 	LastName:    "Sensei",
	// 	Address:     "Tebet",
	// 	Email:       "edo.sensei@enigmacamp.com",
	// 	PhoneNumber: "08219292921",
	// 	Bod:         time.Date(1988, 1, 1, 0, 0, 0, 0, time.UTC),
	// 	UserCredential: model.UserCredential{
	// 		UserName: "edo.sensei@enigmacamp.com",
	// 		Password: password,
	// 		IsActive: false,
	// 	},
	// 	Vehicles: []model.Vehicle{
	// 		{
	// 			BrandID:        "69050e77-e897-4325-9b9e-7480c1f7846f",
	// 			Model:          "Accord",
	// 			ProductionYear: 2022,
	// 			Color:          "Hitam",
	// 			IsAutomatic:    true,
	// 			Stock:          4,
	// 			SalePrice:      700000000,
	// 			Status:         "bekas",
	// 		},
	// 		{
	// 			BrandID:        "69050e77-e897-4325-9b9e-7480c1f7846f",
	// 			Model:          "HR-V",
	// 			ProductionYear: 2023,
	// 			Color:          "Putih",
	// 			IsAutomatic:    true,
	// 			Stock:          3,
	// 			SalePrice:      450000000,
	// 			Status:         "baru",
	// 		},
	// 	},
	// }

	// db.Create(&newCustomerVehicle)

	// Assocuation Mode
	// password, err := utils.HashPassword("P@ssword")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// newCustomer := model.Customer{
	// 	FirstName:   "Suci",
	// 	LastName:    "Dalam Debu",
	// 	Address:     "Surabaya",
	// 	Email:       "suci.dalam@gmail.com",
	// 	PhoneNumber: "028292929",
	// 	Bod:         time.Date(1998, 1, 1, 0, 0, 0, 0, time.UTC),
	// 	UserCredential: model.UserCredential{
	// 		UserName: "suci.dalam@gmail.com",
	// 		Password: password,
	// 		IsActive: false,
	// 	},
	// }
	// db.Create(&newCustomer)

	// // Kedua kita cari vehicle yang mau ditambah
	// var vehicle model.Vehicle
	// if err := db.Where("id=?", "ddba8c27-49b9-4d71-8233-2ebf1de56551").First(&vehicle).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// // Kemudian kita lakukan append
	// if err := db.Model(&vehicle).Association("Customers").Append(&newCustomer); err != nil {
	// 	fmt.Println(err)
	// }

	// Menambahkan dengan customer dan kendaraan yang sudah ada, misalnya SUCI beli mobil baru lagi.
	// var vehicle model.Vehicle
	// if err := db.Debug().Where("id=?", "82a068e3-2367-4db0-b551-731e6bfd4e3a").First(&vehicle).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// var customer model.Customer
	// if err := db.Debug().Where("id=?", "8114e6c6-81b5-4896-8dff-ff67d7ae12d2").First(&customer).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// if err := db.Debug().Model(&vehicle).Association("Customers").Append(&customer); err != nil {
	// 	fmt.Println(err)
	// }

	// Customer menghapus kendaraan nya misalnya
	// var vehicle model.Vehicle
	// if err := db.Debug().Where("id=?", "82a068e3-2367-4db0-b551-731e6bfd4e3a").First(&vehicle).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// var customer model.Customer
	// if err := db.Debug().Preload("UserCredential").Where("id=?", "8114e6c6-81b5-4896-8dff-ff67d7ae12d2").First(&customer).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// if err := db.Debug().Model(&vehicle).Association("Customers").Delete(&customer); err != nil {
	// 	fmt.Println(err)
	// }

	// Customer mengupdate kendaraan yang tersimpan
	// Kendaraan lama:

	// var oldVehicle model.Vehicle
	// if err := db.Debug().Where("id=?", "ddba8c27-49b9-4d71-8233-2ebf1de56551").First(&oldVehicle).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// var newVehicle model.Vehicle
	// if err := db.Debug().Where("id=?", "82a068e3-2367-4db0-b551-731e6bfd4e3a").First(&newVehicle).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// var customer model.Customer
	// if err := db.Debug().Preload("Vehicles").Where("id=?", "a688e737-39f7-4ec7-893a-65e78478580d").First(&customer).Error; err != nil {
	// 	fmt.Println(err)
	// }

	// // ambil id vehicle lama dari slice customer.Vehicles
	// var oldVehicleID string
	// for _, v := range customer.Vehicles {
	// 	if v.ID == oldVehicle.ID {
	// 		oldVehicleID = v.ID
	// 		break
	// 	}
	// }

	// if oldVehicleID == "" {
	// 	fmt.Println("vehicle not found in customer's vehicles")
	// 	return
	// }

	// if err := db.Debug().Model(&customer).Association("Vehicles").Delete([]model.Vehicle{oldVehicle}); err != nil {
	// 	fmt.Println(err)
	// }

	// // set ID baru pada vehicle yang akan diupdate
	// newVehicle.ID = oldVehicleID

	// if err := db.Debug().Model(&customer).Association("Vehicles").Append([]model.Vehicle{newVehicle}); err != nil {
	// 	fmt.Println(err)
	// }

	var customer model.Customer
	if err := db.Debug().Preload("Vehicles").Where("id=?", "a688e737-39f7-4ec7-893a-65e78478580d").First(&customer).Error; err != nil {
		fmt.Println(err)
	}
	var newVehicle model.Vehicle
	if err := db.Debug().Preload("Customers").Where("id=?", "82a068e3-2367-4db0-b551-731e6bfd4e3a").First(&newVehicle).Error; err != nil {
		fmt.Println(err)
	}
	var oldVehicleID = "ddba8c27-49b9-4d71-8233-2ebf1de56551"
	var newVehicleSlice []model.Vehicle
	for _, cv := range customer.Vehicles {
		if cv.ID != oldVehicleID {
			newVehicleSlice = append(newVehicleSlice, cv)
		}
	}
	newVehicleSlice = append(newVehicleSlice, newVehicle)
	if err := db.Debug().Model(&customer).Association("Vehicles").Replace(newVehicleSlice); err != nil {
		fmt.Println(err)
	}
}
