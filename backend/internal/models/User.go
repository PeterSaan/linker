package models

type User struct {
	ID        uint
    Email     string `gorm:"unique"`
	Password  string
	Chats     []Chat
	Chatrooms []*Chatroom `gorm:"many2many:chatroom_users;"`
	JobTitle  []*JobTitle `gorm:"many2many:user_jobtitles;"`
    Profile   Profile
}
