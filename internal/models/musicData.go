package models

type Track struct {
	ID              uint   `gorm:"primaryKey"`
	ISRC            string `gorm:"unique;not null"`
	Title           string
	SpotifyImageURI string
	Artists         []Artist `gorm:"foreignKey:TrackID"`
}

type Artist struct {
	ID      uint `gorm:"primaryKey"`
	TrackID uint `gorm:"index"`
	Name    string
}

type RequestBody struct {
	isrc string `json:"isrc"`
}
