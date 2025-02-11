package models

type User struct {
	ID        uint
	Email     string
	Password  string
	Profile   Profile
	Chats     []Chat
	Chatrooms []*Chatroom     `gorm:"many2many:chatroom_users;"`
	JobTitle  []*JobTitle     `gorm:"many2many:user_jobtitles;"`
}
