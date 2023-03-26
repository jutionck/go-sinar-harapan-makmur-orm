package usecase

import (
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/repository"
)

type VehicleUseCase interface {
	SearchVehicle(by map[string]interface{}) ([]model.Vehicle, error)
	FindAllVehicle() ([]model.Vehicle, error)
	FindVehicleById(id string) (*model.Vehicle, error)
	SaveVehicle(payload *model.Vehicle) (*model.Vehicle, error)
	DeleteVehicle(id string) error
}

type vehicleUseCase struct {
	repo repository.VehicleRepository
}

// SearchVehicle implements VehicleUseCase
func (v *vehicleUseCase) SearchVehicle(by map[string]interface{}) ([]model.Vehicle, error) {
	return v.repo.Search(by)
}

// FindAllVehicle implements VehicleUseCase
func (v *vehicleUseCase) FindAllVehicle() ([]model.Vehicle, error) {
	return v.repo.List()
}

// FindVehicleById implements VehicleUseCase
func (v *vehicleUseCase) FindVehicleById(id string) (*model.Vehicle, error) {
	return v.repo.Get(id)
}

// SaveVehicle implements VehicleUseCase
func (v *vehicleUseCase) SaveVehicle(payload *model.Vehicle) (*model.Vehicle, error) {
	return v.repo.Save(payload)
}

// DeleteVehicle implements VehicleUseCase
func (v *vehicleUseCase) DeleteVehicle(id string) error {
	return v.repo.Delete(id)
}

func NewVehicleUseCase(repo repository.VehicleRepository) VehicleUseCase {
	return &vehicleUseCase{repo: repo}
}
