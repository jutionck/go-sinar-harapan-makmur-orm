package main

import (
	"fmt"
	"log"

	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/config"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
)

func main() {
	fmt.Println("CRUD Interface")

	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err, "", 102)
		return
	}

	newDbConn, err := config.NewDbConnection(c)
	if err != nil {
		log.Fatal(err)
	}

	db := newDbConn.Conn()

	// db.AutoMigrate(&model.Vehicle{})

	// newVehicle := model.Vehicle{
	// 	Brand:          "Honda",
	// 	Model:          "CR-V",
	// 	ProductionYear: 2021,
	// 	Color:          "Putih",
	// 	IsAutomatic:    true,
	// 	Stock:          3,
	// 	SalePrice:      650000000,
	// 	Status:         "Baru",
	// }

	// Create
	// result := db.Create(&newVehicle)
	// result := db.Save(&newVehicle)
	// result := db.Select("Brand", "Model", "Stock", "SalePrice", "CreatedAt").Create(&newVehicle)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }
	// fmt.Println(newVehicle.ID)

	// Update
	// newVehicle.ID = "cd593d53-7ac1-4c65-915d-cb2830667301"
	// newVehicle.Brand = "Toyota"
	// newVehicle.Model = "New Kijang Inova"
	// newVehicle.SalePrice = 550000000
	// result := db.Save(&newVehicle)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// Update Single Column
	// newVehicle.ID = "cd593d53-7ac1-4c65-915d-cb2830667301"
	// result := db.Model(&newVehicle).Update("model", "Innova Zenix")
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// newVehicle.ID = "cd593d53-7ac1-4c65-915d-cb2830667301"
	// result := db.Debug().Unscoped().Model(&model.Vehicle{}).Where("brand = ?", "Toyota").Update("model", "Innova Zenix MN")
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// Update multiple column
	// newVehicle.ID = "cd593d53-7ac1-4c65-915d-cb2830667301"
	// result := db.Model(&model.Vehicle{}).Where("id = ?", newVehicle.ID).Updates(&newVehicle)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// Delete a Record
	// Akan menghasilkan SOFT Delete
	// newVehicle.ID = "cd593d53-7ac1-4c65-915d-cb2830667301"
	// result := db.Delete(&newVehicle)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// Delete dengan ID
	// result := db.Delete(&model.Vehicle{}, "id =?", "9732307e-552b-4962-8da7-b8f8c65030ef")
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// Delete Permanen
	// result := db.Unscoped().Delete(&model.Vehicle{}, "id =?", "9732307e-552b-4962-8da7-b8f8c65030ef")
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// Select All
	// var vehicles []model.Vehicle
	// result := db.Find(&vehicles)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// for _, v := range vehicles {
	// 	fmt.Println("ID:", v.ID)
	// 	fmt.Println("Brand:", v.Brand)
	// 	fmt.Println("Model:", v.Model)
	// 	fmt.Println("Production Year:", v.ProductionYear)
	// 	fmt.Println("Price:", v.SalePrice)
	// 	fmt.Println()
	// }

	// Jika tidak ditemukan berarti SELECT nya deleted is NULL, unutk mengatasi nya adalah dengan Unscoped()
	// result := db.Unscoped().Find(&vehicles)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// fmt.Println(vehicles)

	// Select Single Row
	// var vehicle model.Vehicle
	// result := db.Unscoped().First(&vehicle, "id = ?", "46742391-0bd7-46e5-b2a3-a0980efcd998")
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// fmt.Println("ID:", vehicle.ID)
	// fmt.Println("Brand:", vehicle.Brand)
	// fmt.Println("Model:", vehicle.Model)
	// fmt.Println("Production Year:", vehicle.ProductionYear)
	// fmt.Println("Price:", vehicle.SalePrice)

	// Refinig Selection
	// Where Clause
	// var vehicles []model.Vehicle
	// result := db.Unscoped().Where("model = ?", "hr-v").Find(&vehicles)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// for _, v := range vehicles {
	// 	fmt.Println("ID:", v.ID)
	// 	fmt.Println("Brand:", v.Brand)
	// 	fmt.Println("Model:", v.Model)
	// 	fmt.Println("Production Year:", v.ProductionYear)
	// 	fmt.Println("Price:", v.SalePrice)
	// 	fmt.Println()
	// }

	// WHERE IN
	// var vehicles []model.Vehicle
	// result := db.Unscoped().Where("model IN ?", []string{"HR-V", "CR-V"}).Find(&vehicles)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// for _, v := range vehicles {
	// 	fmt.Println("ID:", v.ID)
	// 	fmt.Println("Brand:", v.Brand)
	// 	fmt.Println("Model:", v.Model)
	// 	fmt.Println("Production Year:", v.ProductionYear)
	// 	fmt.Println("Price:", v.SalePrice)
	// 	fmt.Println()
	// }

	// LIKE
	// var vehicles []model.Vehicle
	// result := db.Where("model LIKE ?", "%R%").Find(&vehicles)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }
	// for _, v := range vehicles {
	// 	fmt.Println("ID:", v.ID)
	// 	fmt.Println("Brand:", v.Brand)
	// 	fmt.Println("Model:", v.Model)
	// 	fmt.Println("Production Year:", v.ProductionYear)
	// 	fmt.Println("Price:", v.SalePrice)
	// 	fmt.Println()
	// }

	// Paging
	// var vehicles []model.Vehicle
	// page := 1
	// itemPerPage := 2
	// offset := itemPerPage * (page - 1)
	// result := db.Unscoped().Order("created_at").Limit(itemPerPage).Offset(offset).Find(&vehicles)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// }

	// for _, v := range vehicles {
	// 	fmt.Println("ID:", v.ID)
	// 	fmt.Println("Brand:", v.Brand)
	// 	fmt.Println("Model:", v.Model)
	// 	fmt.Println("Production Year:", v.ProductionYear)
	// 	fmt.Println("Price:", v.SalePrice)
	// 	fmt.Println()
	// }

	// Aggregate
	// Menghitung total kendaraan
	// var total int64
	// db.Model(&model.Vehicle{}).Select("COUNT(*) as total").Count(&total)
	// fmt.Println("Total Kendaraan:", total)
	var vehicleBrandCount []VehicleBrandCount
	db.Model(&model.Vehicle{}).Select("brand, COUNT(*) as total").Group("brand").Find(&vehicleBrandCount)
	for _, v := range vehicleBrandCount {
		fmt.Println("Brand:", v.Brand)
		fmt.Println("Total:", v.Total)
		fmt.Println()
	}

	// SQL Builder (Raw Query)
}

type VehicleBrandCount struct {
	Brand string
	Total int
}
