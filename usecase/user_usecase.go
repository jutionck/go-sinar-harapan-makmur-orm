package usecase

import (
	"fmt"

	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/repository"
)

type UserUseCase interface {
	BaseUseCase[model.UserCredential]
	FindUserByUsername(username string) (*model.UserCredential, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func (u *userUseCase) DeleteData(id string) error {
	user, err := u.FindById(id)
	if err != nil {
		return fmt.Errorf("User with ID %s not found!", id)
	}
	return u.repo.Delete(user.ID)
}

func (u *userUseCase) FindAll() ([]model.UserCredential, error) {
	return u.repo.List()
}

func (u *userUseCase) FindById(id string) (*model.UserCredential, error) {
	user, err := u.repo.Get(id)
	if err != nil {
		return nil, fmt.Errorf("User with ID %s not found!", id)
	}
	return user, nil
}

func (u *userUseCase) SaveData(payload *model.UserCredential) error {
	if payload.ID != "" {
		_, err := u.FindById(payload.ID)
		if err != nil {
			return fmt.Errorf("User with ID %s not found!", payload.ID)
		}
	}
	return u.repo.Save(payload)
}

func (u *userUseCase) SearchBy(by map[string]interface{}) ([]model.UserCredential, error) {
	users, err := u.repo.Search(by)
	if err != nil {
		return nil, fmt.Errorf("Data not found")
	}
	return users, nil
}

func (u *userUseCase) FindUserByUsername(username string) (*model.UserCredential, error) {
	user, err := u.repo.GetByUsername(username)
	if err != nil {
		return nil, fmt.Errorf("User with username %s not found!", username)
	}
	return user, nil
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}
