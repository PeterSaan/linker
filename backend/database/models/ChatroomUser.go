package models

type ChatroomUser struct {
	ID         int64 `gorm:"primaryKey;autoIncrement:false"`
	ChatroomID int64 `gorm:"column:chatroom_id"`
	UserID     int64 `gorm:"column:user_id"`
}
