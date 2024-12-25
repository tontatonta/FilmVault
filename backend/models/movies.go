package models

type Movies struct {
	MovieID int    `gorm:"primaryKey"`  // 主キーとしてMovieID
	Title   string `gorm:"size:255"`    // 映画のタイトル
	Image   string `gorm:"size:255"`    // 映画の画像URL
	GenreID int    `gorm:"index"`       // ジャンルID（インデックス）
	UserMovie []UserMovie `gorm:"foreignKey:MovieID"`  // UserMovieとの関連付け（外部キー: MovieID）
	Genre   Genre  `gorm:"foreignKey:GenreID"` // Genreモデルとの関連付け（外部キー: GenreID）
}
