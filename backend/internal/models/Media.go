package models

type Media struct {
	ID        uint
	URL       string
	ProfileID uint
	Profile   Profile
}
