package models

type Chat struct {
	ID         uint		`gorm:"primaryKey;autoIncrement:false"`
	Body       string	`gorm:"column:body"`
	UserID     uint		`gorm:"column:user_id"`
	ChatroomID uint		`gorm:"column:chatroom_id"`
}
