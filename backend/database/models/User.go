package models

type User struct {
	ID        int64          `gorm:"primaryKey;autoIncrement:false"`
	Email     int64          `gorm:"column:email"`
	Password  int64          `gorm:"column:password"`
	Profile   Profile        `gorm:"foreignKey:ID;references:UserID"`
	Chats     []Chat         `gorm:"foreignKey:UserID"`
	Chatrooms []ChatroomUser `gorm:"foreignKey:UserID"`
	JobTitle  UserJobTitle   `gorm:"foreignKey:UserID"`
}
