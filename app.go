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

	// dbConn.Migrate(&model.Brand{}, &model.Vehicle{}, &model.Customer{}, &model.UserCredential{})

	// create brand (batch)
	// brands := []model.Brand{
	// 	{Name: "Honda"},
	// 	{Name: "Toyota"},
	// 	{Name: "Daihatsu"},
	// 	{Name: "Mazda"},
	// 	{Name: "Mitshubisi"},
	// 	{Name: "Suzuki"},
	// }
	// result := db.Create(&brands).Error
	// if result != nil {
	// 	fmt.Println(result)
	// }

	// Create vehicle use brand
	// get brand first
	// var brand model.Brand
	// if result := db.First(&brand, "name =?", "Honda").Error; result != nil {
	// 	fmt.Println(result)
	// }

	// vehicle := model.Vehicle{
	// 	BrandID:        brand.ID,
	// 	Model:          "Jazz",
	// 	ProductionYear: 2022,
	// 	Color:          "Putih",
	// 	IsAutomatic:    true,
	// 	Stock:          10,
	// 	SalePrice:      240000000,
	// 	Status:         "Bekas",
	// }

	// if result := db.Create(&vehicle).Error; result != nil {
	// 	fmt.Println(result)
	// }

	// Create brand with Vehicle
	// brand := model.Brand{
	// 	Name: "Honda",
	// 	Vehicle: []model.Vehicle{
	// 		{
	// 			Model:          "Jazz",
	// 			ProductionYear: 2023,
	// 			Color:          "Kuning",
	// 			IsAutomatic:    true,
	// 			Stock:          3,
	// 			SalePrice:      350000000,
	// 			Status:         "Baru",
	// 		},
	// 		{
	// 			Model:          "City",
	// 			ProductionYear: 2023,
	// 			Color:          "Hitam",
	// 			IsAutomatic:    false,
	// 			Stock:          5,
	// 			SalePrice:      450000000,
	// 			Status:         "Baru",
	// 		},
	// 	},
	// }
	// result := db.Create(&brand).Error
	// if result != nil {
	// 	fmt.Println(result)
	// }

	// List Brand With Vehicle
	var brands []model.Brand
	db.Preload("Vehicle").Find(&brands)

	for _, brand := range brands {
		fmt.Printf("Brand: %s\n", brand.Name)
		for _, vehicle := range brand.Vehicle {
			fmt.Println("Model:", vehicle.Model)
			fmt.Println("Sale Price:", vehicle.Model)
			fmt.Println("Color:", vehicle.Model)
			fmt.Println()
		}
	}

	// Get Brand With Vehicle
	// var brand model.Brand
	// db.Preload("Vehicle").First(&brand, "name = ?", "Toyota")
	// for _, vehicle := range brand.Vehicle {
	// 	fmt.Println("Brand:", brand.Name)
	// 	fmt.Println("Model:", vehicle.Model)
	// 	fmt.Println("Sale Price:", vehicle.Model)
	// 	fmt.Println("Color:", vehicle.Model)
	// 	fmt.Println()
	// }

}
