package models

type ChatroomUser struct {
	ID         uint	`gorm:"primaryKey;autoIncrement:false"`
	ChatroomID uint	`gorm:"column:chatroom_id"`
	UserID     uint	`gorm:"column:user_id"`
}
