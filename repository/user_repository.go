package repository

import (
	"github.com/jutionck/golang-db-sinar-harapan-makmur-orm/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	BaseRepository[model.UserCredential]
	GetByUsername(username string) (*model.UserCredential, error)
}

type userRepository struct {
	db *gorm.DB
}

func (u *userRepository) Search(by map[string]interface{}) ([]model.UserCredential, error) {
	var users []model.UserCredential
	result := u.db.Where(by).Find(&users).Error
	if result != nil {
		return nil, result
	}
	return users, nil
}

func (u *userRepository) List() ([]model.UserCredential, error) {
	var users []model.UserCredential
	result := u.db.Find(&users).Error
	if result != nil {
		return nil, result
	}
	return users, nil
}

func (u *userRepository) Get(id string) (*model.UserCredential, error) {
	var user model.UserCredential
	result := u.db.First(&user, "id=?", id).Error
	if result != nil {
		return nil, result
	}
	return &user, nil
}

func (u *userRepository) Save(payload *model.UserCredential) error {
	return u.db.Save(payload).Error
}

func (u *userRepository) Delete(id string) error {
	return u.db.Delete(&model.UserCredential{}, "id=?", id).Error
}

func (u *userRepository) GetByUsername(username string) (*model.UserCredential, error) {
	var userCredential model.UserCredential
	result := u.db.First(&userCredential, "username=? and is_active = true", username).Error
	if result != nil {
		return nil, result
	}
	return &userCredential, nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}
