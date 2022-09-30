package models

import "time"

type User struct {
	Id        int               `json:"id" gorm:"primary_key:auto_increment"`
	Name      string            `json:"name" gorm:"type: varchar(255)"`
	Email     string            `json:"email" gorm:"type: varchar(255)"`
	Password  string            `json:"password" gorm:"type: varchar(255)"`
	Profile   ProfileRelation   `json:"profile"`
	Product   []ProductRelation `json:"products"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

type UserRelation struct {
	Id    int    `json:"-"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (UserRelation) TableName() string { return "users" }
