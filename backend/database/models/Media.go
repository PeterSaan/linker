package models

type Media struct {
	ID        int64 `gorm:"primaryKey;autoIncrement:false"`
	URL       int64 `gorm:"column:url"`
	ProfileID int64 `gorm:"column:profile_id"`
}
