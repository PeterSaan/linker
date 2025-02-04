package models

type Profile struct {
	ID          int64
	UserID      int64
	Name        int64
	Description int64
	Avatar      int64
	Type        int64
	Media       []Media
}
