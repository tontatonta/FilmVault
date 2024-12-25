package models

type UserMovie struct {
	UserID  int    `gorm:"primaryKey;autoIncrement:false"` // 複合主キーの一部として設定
	MovieID int    `gorm:"primaryKey;autoIncrement:false"` // 複合主キーの一部として設定
	User    Users  `gorm:"foreignKey:UserID"`
	Movie   Movies  `gorm:"foreignKey:MovieID"`
}
