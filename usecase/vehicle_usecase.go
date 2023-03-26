package usecase

import (
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/repository"
)

type VehicleUseCase interface {
	BaseUseCase[model.Vehicle]
}

type vehicleUseCase struct {
	repo repository.VehicleRepository
}

func (v *vehicleUseCase) SearchBy(by map[string]interface{}) ([]model.Vehicle, error) {
	return v.repo.Search(by)
}

func (v *vehicleUseCase) FindAll() ([]model.Vehicle, error) {
	return v.repo.List()
}

func (v *vehicleUseCase) FindById(id string) (*model.Vehicle, error) {
	return v.repo.Get(id)
}

func (v *vehicleUseCase) SaveData(payload *model.Vehicle) error {
	return v.repo.Save(payload)
}

func (v *vehicleUseCase) DeleteData(id string) error {
	return v.repo.Delete(id)
}

func NewVehicleUseCase(repo repository.VehicleRepository) VehicleUseCase {
	return &vehicleUseCase{repo: repo}
}
