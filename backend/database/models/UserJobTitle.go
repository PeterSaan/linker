package models

type UserJobTitle struct {
	UserID     int64 `gorm:"primaryKey;autoIncrement:false"`
	JobTitleID int64 `gorm:"column:job_title_id"`
}
