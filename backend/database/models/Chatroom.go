package models

type Chatroom struct {
	ID    int64          `gorm:"primaryKey;autoIncrement:false"`
	Name  int64          `gorm:"column:name"`
	Chats []Chat         `gorm:"foreignKey:ChatroomID"`
	Users []ChatroomUser `gorm:"foreignKey:ChatroomID"`
}
