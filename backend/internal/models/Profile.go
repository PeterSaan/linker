package models

type Profile struct {
	ID          uint
	UserID      uint `gorm:"unique"`
	FirstName   string
	LastName    string
	Location   	string
	Description string
	Avatar      string
	Type        string
	Media       []Media `gorm:"many2many:profile_media;"`
}
