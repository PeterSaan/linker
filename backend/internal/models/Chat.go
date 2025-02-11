package models

type Chat struct {
	ID         uint
	Body       string
	UserID     uint
	User       User
	ChatroomID uint
	Chatroom   Chatroom
}
