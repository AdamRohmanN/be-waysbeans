package models

import "time"

type Transaction struct {
	Id        int             `json:"id" gorm:"primary_key:auto_increment"`
	ProductId int             `json:"product_id"`
	Product   ProductRelation `json:"product"`
	BuyerId   int             `json:"buyer_id"`
	Buyer     UserRelation    `json:"buyer"`
	SellerId  int             `json:"seller_id"`
	Seller    UserRelation    `json:"seller"`
	Price     int             `json:"price" gorm:"type: int"`
	Status    string          `json:"status" gorm:"type: varchar(255)"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
