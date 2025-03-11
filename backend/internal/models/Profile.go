package models

type Profile struct {
	ID          uint
	UserID      uint `gorm:"unique"`
	Name        string
	Description string
	Avatar      string
	Type        string
	Media       []Media `gorm:"many2many:profile_media;"`
}
