package models

type Chatroom struct {
	ID    uint						
	Name  string					
	Chats []Chat					
	Users  []*User `gorm:"many2many:chatroom_users;"`
}
