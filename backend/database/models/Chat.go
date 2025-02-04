package models

type Chat struct {
	ID         int64 `gorm:"primaryKey;autoIncrement:false"`
	Body       int64 `gorm:"column:body"`
	UserID     int64 `gorm:"column:user_id"`
	ChatroomID int64 `gorm:"column:chatroom_id"`
}
