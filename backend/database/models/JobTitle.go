type JobTitle struct {
	ID           int64        `gorm:"primaryKey;autoIncrement:false"`
	Name         int64        `gorm:"column:name"`
	UserJobTitle UserJobTitle `gorm:"foreignKey:JobTitleID"`
}
