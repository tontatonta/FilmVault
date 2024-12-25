package models

type Users struct {
	UserID   int        `gorm:"primaryKey"` // 主キーとしてUserID
	Email    string     `gorm:"unique;not null"` // メール（ユニーク、NULL不可）
	Password string     `gorm:"size:255"` // パスワード（サイズ指定）
	UserMovie []UserMovie `gorm:"foreignKey:UserID"` // UserMovieとの関連付け（外部キー: UserID）
}
