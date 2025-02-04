package models

type Profile struct {
	ID          int64   `gorm:"primaryKey;autoIncrement:false"`
	UserID      int64   `gorm:"column:user_id"`
	Name        int64   `gorm:"column:name"`
	Description int64   `gorm:"column:description"`
	Avatar      int64   `gorm:"column:avatar"`
	Type        int64   `gorm:"column:type"`
	Media       []Media `gorm:"foreignKey:ProfileID"`
}
