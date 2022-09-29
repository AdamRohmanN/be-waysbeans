package repositories

import (
	"time"
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
	err := r.db.Raw("SELECT * FROM users").Scan(&users).Error

	return users, err
}

func (r *repository) GetUser(Id int) (models.User, error) {
	var user models.User
	err := r.db.Raw("SELECT * FROM users WHERE id=?", Id).Scan(&user).Error

	return user, err
}

func (r *repository) CreateUser(user models.User) (models.User, error) {
	err := r.db.Exec("INSERT INTO users(name,email,password,created_at,updated_at) VALUES (?,?,?,?,?)",
		user.Name, user.Email, user.Password, time.Now(), time.Now()).Error

	return user, err
}

func (r *repository) UpdateUser(user models.User, Id int) (models.User, error) {
	err := r.db.Raw("UPDATE users SET name=?, email=?, password=? WHERE id=?",
		user.Name, user.Email, user.Password, Id).Scan(&user).Error

	return user, err
}

func (r *repository) DeleteUser(user models.User, Id int) (models.User, error) {
	err := r.db.Raw("DELETE FROM users WHERE id=?", Id).Scan(&user).Error

	return user, err
}
