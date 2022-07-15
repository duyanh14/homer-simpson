package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"column:username;unique" bson:"username"`
	Password string `gorm:"column:password" bson:"password"`
	Phone    string `gorm:"column:phone" bson:"phone"`
	Email    string `gorm:"column:email" bson:"email"`
}

func (u User) Table() string {
	return "users"
}
