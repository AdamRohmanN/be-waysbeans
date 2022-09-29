package userdto

import "waysbeans/models"

type UserResponse struct {
	Name    string                 `json:"name"`
	Email   string                 `json:"email"`
	Profile models.ProfileRelation `json:"profile"`
}
