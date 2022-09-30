package models

import "time"

type Product struct {
	Id         int          `json:"id" gorm:"primary_key:auto_increment"`
	Image      string       `json:"image" gorm:"type: varchar(255)"`
	Name       string       `json:"name" gorm:"type: varchar(255)"`
	Desc       string       `json:"desc" gorm:"type: varchar(255)"`
	Price      int          `json:"price" gorm:"type: int"`
	Stock      int          `json:"stock" gorm:"type: int"`
	UserId     int          `json:"user_id"`
	User       UserRelation `json:"user"`
	CategoryId []int        `json:"category_id"`
	Category   []Category   `json:"categories" gorm:"many2many:product_categories"`
	CreatedAt  time.Time    `json:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at"`
}

type ProductRelation struct {
	Id         int        `json:"id"`
	Image      string     `json:"image"`
	Name       string     `json:"name"`
	Desc       string     `json:"desc"`
	Price      int        `json:"price"`
	Stock      int        `json:"stock"`
	UserId     int        `json:"-"`
	CategoryId []int      `json:"-"`
	Category   []Category `json:"categories"`
}

func (ProductRelation) TableName() string { return "products" }
