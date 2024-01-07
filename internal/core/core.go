package core

import (
	model "CodingChallenge/internal/models"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"net/http"
)

type DBContainer struct {
	DB            *gorm.DB
}

// NewContainer creates a new Container with dependencies.
func NewDBContainer() *DBContainer {

	dsn := "user=postgres password=password dbname=mydatabase sslmode=disable host=localhost"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Track{}, &model.Artist{})

	return &DBContainer{
		DB:            db,
	}
}

func CreateTrack(c *gin.Context, container *DBContainer) {
	// Parse JSON request
	var requestBody map[string]string
	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	isrc := requestBody["isrc"]

	// Fetch track metadata from Spotify
	track, err := fetchTrackMetadata(isrc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch track metadata from Spotify"})
		return
	}

	// Store track in the database
	rows := container.DB.Create(&track)
	if rows.Error != nil {
		fmt.Print("error while creating entry in  database")
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Track created successfully"})
}

func GetTrackByISRC(c *gin.Context, container *DBContainer) {
	isrc := c.Param("isrc")

	var track model.Track
	if err := container.DB.Preload("Artists").Where("isrc = ?", isrc).First(&track).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Track not found"})
		return
	}

	c.JSON(http.StatusOK, track)
}

func GetTracksByArtist(c *gin.Context, container *DBContainer) {
	artistName := c.Param("artist")

	var tracks []model.Track
	if err := container.DB.
		Preload("Artists").
		Select("tracks.id, tracks.isrc, tracks.title, tracks.spotify_image_uri").
		Joins("JOIN artists ON artists.track_id = tracks.id").
		Where("artists.name LIKE ?", fmt.Sprintf("%%%s%%", artistName)).
		Find(&tracks).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tracks not found"})
		return
	}

	c.JSON(http.StatusOK, tracks)
}

func fetchTrackMetadata(isrc string) (model.Track, error) {
	// Spotify API configuration
	config := &clientcredentials.Config{
		ClientID:     "fb26292c91fc4c4c9519d02e063f37be",
		ClientSecret: "59ac994379fe40fb889044b50f3173ef",
		TokenURL:     spotify.TokenURL,
	}

	token, err := config.Token(context.Background())
	if err != nil {
		return model.Track{}, err
	}

	// Create Spotify client
	auth := spotify.NewAuthenticator(token.AccessToken)
	client := auth.NewClient(token)

	// Search for track using ISRC
	result, err := client.Search(fmt.Sprintf("isrc:%s", isrc), spotify.SearchTypeTrack)
	if err != nil {
		return model.Track{}, err
	}

	// Get the first track (highest popularity)
	firstTrack := result.Tracks.Tracks[0]

	// Extract track metadata
	track := model.Track{
		ISRC:            isrc,
		Title:           firstTrack.Name,
		SpotifyImageURI: firstTrack.Album.Images[0].URL,
	}

	// Extract artist names
	for _, artist := range firstTrack.Artists {
		track.Artists = append(track.Artists, model.Artist{Name: artist.Name})
	}

	return track, nil
}
