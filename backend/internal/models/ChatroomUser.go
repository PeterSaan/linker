package models

type ChatroomUser struct {
	ID         uint
	ChatroomID uint
	Chatroom   Chatroom
	UserID     uint
	User       []*User `gorm:"many2many:chatroom_users;"`
}
