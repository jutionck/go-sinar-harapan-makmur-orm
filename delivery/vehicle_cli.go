package delivery

import (
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/usecase"
)

type VehicleCli struct {
	vehicleUseCase usecase.VehicleUseCase
}

// Service CRUD disini
func (v *VehicleCli) registerVehicle() {
	newVehicle := &model.Vehicle{
		Brand:          "Honda",
		Model:          "Jazz",
		ProductionYear: 2022,
		SalePrice:      550000000,
		Color:          "Putih",
		IsAutomatic:    true,
		Stock:          5,
		Status:         "Baru",
	}
	v.vehicleUseCase.SaveVehicle(newVehicle)
}

func NewVehicleCli(vehicleUseCase usecase.VehicleUseCase, menu string) *VehicleCli {
	cli := VehicleCli{vehicleUseCase: vehicleUseCase}
	switch menu {
	case "VH1":
		cli.registerVehicle()
	case "VH2":
		cli.registerVehicle()
	case "VH3":
		cli.registerVehicle()
	case "VH4":
		cli.registerVehicle()
	case "VH15":
		cli.registerVehicle()
	}
	return &cli
}
