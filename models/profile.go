package models

import "time"

type Profile struct {
	Id        int          `json:"id" gorm:"primary_key:auto_increment"`
	Image     string       `json:"image" gorm:"type: varchar(255)"`
	Address   string       `json:"address" gorm:"type: varchar(255)"`
	Postcode  string       `json:"postcode" gorm:"type: varchar(255)"`
	UserId    int          `json:"user_id"`
	User      UserRelation `json:"user"`
	UpdatedAt time.Time    `json:"updated_at"`
}

type ProfileRelation struct {
	Image    string `json:"image"`
	Address  string `json:"address"`
	Postcode string `json:"postcode"`
	UserId   int    `json:"-"`
}

func (ProfileRelation) TableName() string { return "profiles" }
