package main

import (
	"github.com/gin-gonic/gin"
    core "CodingChallenge/internal/core"
)

// @title LTIMindTree Coding Exercise
// @version 1.0
// @description API Coding Challenge Summary For LTIMindTree
// @BasePath /api
func main() {

    dbContainer := core.NewDBContainer()

	router := gin.Default()


	// Create track endpoint
    // @Summary Create a new track
    // @Tags tracks
    // @Accept json
    // @Produce json
    // @Param track body TrackRequest true "Track data"
    // @Success 201 "{'message':'Track created successfully'}"
    // @Router /tracks [post]
	router.POST("/api/tracks", func(c *gin.Context) {
		core.CreateTrack(c, dbContainer)
	})

	// Get track by ISRC endpoint
    // @Summary Get a track by ISRC
    // @Tags tracks
    // @Accept json
    // @Produce json
    // @Param isrc path string true "ISRC of the track"
    // @Success 200 {object} Track
    // @Router /tracks/{isrc} [get]
	router.GET("/api/tracks/:isrc", func(c *gin.Context) {
		core.GetTrackByISRC(c, dbContainer)
	})

	// Get tracks by artist endpoint
    // @Summary Get tracks by artist
    // @Tags tracks
    // @Accept json
    // @Produce json
    // @Param artist path string true "Artist name"
    // @Success 200 {array} Track
    // @Router /tracks/by-artist/{artist} [get]
	router.GET("/api/tracks/by-artist/:artist", func(c *gin.Context) {
		core.GetTracksByArtist(c, dbContainer)
	})

	router.Run(":8080")
}
