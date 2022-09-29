package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	GetProfile(Id int) (models.Profile, error)
}

func RepositoryProfile(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetProfile(Id int) (models.Profile, error) {
	var profile models.Profile
	err := r.db.Preload("User").First(&profile, Id).Error

	return profile, err
}
