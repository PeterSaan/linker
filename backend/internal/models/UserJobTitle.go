package models

type UserJobTitle struct {
	UserID     uint	`gorm:"primaryKey;autoIncrement:false"`
	JobTitleID uint	`gorm:"column:job_title_id"`
}
