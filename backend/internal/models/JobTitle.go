package models

type JobTitle struct {
	ID           uint					`gorm:"primaryKey;autoIncrement:false"`
	Name         string				`gorm:"column:name"`
	UserJobTitle UserJobTitle	`gorm:"foreignKey:JobTitleID"`
}
