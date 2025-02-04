package models

type User struct {
	ID        int64
	Email     int64
	Password  int64
	Profile   Profile
	Chats     []Chat
	Chatrooms []ChatroomUser
	JobTitle  UserJobTitle
}
