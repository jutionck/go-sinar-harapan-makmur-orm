package main

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/config"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/repository"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/usecase"
)

func main() {
	c, err := config.NewConfig()
	if err != nil {
		fmt.Println(err)
	}

	dbConn, _ := config.NewDbConnection(c)
	db := dbConn.Conn()

	// Vehicle
	vehicleRepo := repository.NewVehicleRepository(db)
	vehicleUseCase := usecase.NewVehicleUseCase(vehicleRepo)

	// newVehicle := model.Vehicle{
	// 	Brand:          "Honda",
	// 	Model:          "HR-v",
	// 	ProductionYear: 2022,
	// 	Color:          "Silver",
	// 	IsAutomatic:    true,
	// 	Stock:          5,
	// 	SalePrice:      650000000,
	// 	Status:         "Baru",
	// }

	// if err := vehicleUseCase.SaveData(&newVehicle); err != nil {
	// 	fmt.Println(err)
	// }

	// Search
	search := map[string]interface{}{
		"brand": "Honda",
		"model": "HR-v",
	}
	vehicles, err := vehicleUseCase.SearchBy(search)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range vehicles {
		fmt.Println("ID:", v.ID)
		fmt.Println("Brand:", v.Brand)
		fmt.Println("Model:", v.Model)
		fmt.Println("Production Year:", v.ProductionYear)
		fmt.Println("Price:", v.SalePrice)
		fmt.Println()
	}

}
