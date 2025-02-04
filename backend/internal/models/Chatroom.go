package models

type Chatroom struct {
	ID    int64
	Name  int64
	Chats []Chat
	Users []ChatroomUser
}
