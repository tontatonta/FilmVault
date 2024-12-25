package models

type Genre struct {
	GenreID   int      `gorm:"primaryKey"`   // 主キーとしてGenreID
	GenreName string   `gorm:"size:255"`     // ジャンル名
	Movies    []Movies  `gorm:"foreignKey:GenreID"` // Movieとの関連付け（外部キー: GenreID）
}
