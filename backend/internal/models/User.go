package models

type User struct {
	ID        uint						`gorm:"primaryKey;autoIncrement:false"`
	Email     string					`gorm:"column:email"`
	Password  string					`gorm:"column:password"`
	Profile   Profile					`gorm:"foreignKey:ID;references:UserID"`
	Chats     []Chat					`gorm:"foreignKey:UserID"`
	Chatrooms []ChatroomUser	`gorm:"foreignKey:UserID"`
	JobTitle  UserJobTitle		`gorm:"foreignKey:UserID"`
}
