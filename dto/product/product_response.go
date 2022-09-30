package productdto

import "waysbeans/models"

type ProductResponse struct {
	Id       int                 `json:"id"`
	Image    string              `json:"image"`
	Name     string              `json:"name"`
	Desc     string              `json:"desc"`
	Price    int                 `json:"price"`
	Stock    int                 `json:"stock"`
	User     models.UserRelation `json:"user"`
	Category []models.Category   `json:"categories"`
}
