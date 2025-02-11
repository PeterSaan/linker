package models

type Chatroom struct {
	ID    uint						
	Name  string					
	Chats []Chat					
	Users []ChatroomUser	
}
