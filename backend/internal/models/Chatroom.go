package models

type Chatroom struct {
	ID    uint						`gorm:"primaryKey;autoIncrement:false"`
	Name  string					`gorm:"column:name"`
	Chats []Chat					`gorm:"foreignKey:ChatroomID"`
	Users []ChatroomUser	`gorm:"foreignKey:ChatroomID"`
}
