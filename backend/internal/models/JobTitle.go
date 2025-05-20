package models

type JobTitle struct {
	ID   uint
	Name string
	User []*User `gorm:"many2many:user_jobtitles;"`
}
