package models

type Track struct {
	ID              uint   `gorm:"primaryKey"`
	ISRC            string `gorm:"unique;not null"`
	Title           string
	SpotifyImageURI string
	Artists         []Artist `gorm:"foreignKey:TrackID"` // Relationship: One Track has many Artists
}

type Artist struct {
	ID      uint `gorm:"primaryKey"`
	TrackID uint `gorm:"index"` // Foreign key
	Name    string
}

type RequestBody struct {
	isrc string `json:"isrc"`
}
