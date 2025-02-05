package models

type Profile struct {
	ID          uint		`gorm:"primaryKey;autoIncrement:false"`
	UserID      uint		`gorm:"column:user_id"`
	Name        string	`gorm:"column:name"`
	Description string	`gorm:"column:description"`
	Avatar      string	`gorm:"column:avatar"`
	Type        string	`gorm:"column:type"`
	Media       []Media	`gorm:"foreignKey:ProfileID"`
}
