package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Posts    []Post `gorm:"foreignKey:UserID"`
}

type Post struct {
	gorm.Model
	Title   string
	Content string
	UserID  uint
	Author  User `gorm:"foreignKey:UserID" json:"author"`
}
