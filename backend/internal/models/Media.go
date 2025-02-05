package models

type Media struct {
	ID        uint	`gorm:"primaryKey;autoIncrement:false"`
	URL       string	`gorm:"column:url"`
	ProfileID uint	`gorm:"column:profile_id"`
}
