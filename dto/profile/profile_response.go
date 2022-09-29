package profiledto

import "waysbeans/models"

type ProfileResponse struct {
	Image    string              `json:"image"`
	Address  string              `json:"address"`
	Postcode string              `json:"postcode"`
	User     models.UserRelation `json:"user"`
}
