package models

type Profile struct {
	ID          uint
    User        User
	UserID      uint
	Name        string
	Description string
	Avatar      string
	Type        string
    Media       []Media `gorm:"many2many:profile_media;"`
}
