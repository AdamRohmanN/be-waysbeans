package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	FindUsers() ([]models.User, error)
	GetUser(Id int) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User, Id int) (models.User, error)
	DeleteUser(user models.User, Id int) (models.User, error)
}

func RepositoryUser(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindUsers() ([]models.User, error) {
	var users []models.User
	err := r.db.Preload("Profile").Find(&users).Error

	return users, err
}

func (r *repository) GetUser(Id int) (models.User, error) {
	var user models.User
	err := r.db.Preload("Profile").First(&user, Id).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User, Id int) (models.User, error) {
	err := r.db.Save(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User, Id int) (models.User, error) {
	err := r.db.Delete(&user).Error

	return user, err
}
