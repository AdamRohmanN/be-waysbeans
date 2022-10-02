package repositories

import (
	"waysbeans/models"

	"gorm.io/gorm"
)

type ProfileRepository interface {
	GetProfile(Id int) (models.Profile, error)
	CreateProfile(profile models.Profile) (models.Profile, error)
	UpdateProfile(profile models.Profile, Id int) (models.Profile, error)
}

func RepositoryProfile(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetProfile(Id int) (models.Profile, error) {
	var profile models.Profile
	err := r.db.Preload("User").First(&profile, Id).Error

	return profile, err
}

func (r *repository) CreateProfile(profile models.Profile) (models.Profile, error) {
	err := r.db.Create(&profile).Error

	return profile, err
}

func (r *repository) UpdateProfile(profile models.Profile, Id int) (models.Profile, error) {
	err := r.db.Debug().Save(&profile).Error

	return profile, err
}
