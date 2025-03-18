package models

import (
	"strings"

	"gorm.io/gorm"
)

type User struct {
	ID        uint
	Email     string `gorm:"unique"`
	Password  string
	Chats     []Chat
	Chatrooms []*Chatroom `gorm:"many2many:chatroom_users;"`
	JobTitle  []*JobTitle `gorm:"many2many:user_jobtitles;"`
	Profile   Profile
}

func (user *User) BeforeCreate(tx *gorm.DB) (err error) {
    if user != nil {
        user.Email = strings.ToLower(user.Email)
    }
	return
}
