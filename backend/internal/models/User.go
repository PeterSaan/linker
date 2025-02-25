package models

type User struct {
	ID        uint   `json:"user_id"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Chats     []Chat
	Chatrooms []*Chatroom `gorm:"many2many:chatroom_users;"`
	JobTitle  []*JobTitle `gorm:"many2many:user_jobtitles;"`
}
