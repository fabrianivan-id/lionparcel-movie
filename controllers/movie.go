package controllers

import (
	"net/http"
	"strconv"

	"movie-festival-app/config"
	"movie-festival-app/models"

	"github.com/gin-gonic/gin"
)

func CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create movie"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Movie created successfully", "movie": movie})
}

func VoteMovie(c *gin.Context) {
	userID := c.GetString("user_id") // Extract from JWT
	movieID := c.Param("id")

	vote := models.Vote{UserID: userID, MovieID: movieID}
	if err := config.DB.Create(&vote).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to vote for the movie"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Voted successfully"})
}

func UnvoteMovie(c *gin.Context) {
	userID := c.GetString("user_id")
	movieID := c.Param("id")

	if err := config.DB.Where("user_id = ? AND movie_id = ?", userID, movieID).Delete(&models.Vote{}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to unvote the movie"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Unvoted successfully"})
}

func UpdateMovie(c *gin.Context) {
	movieID := c.Param("id")
	var movieUpdates models.Movie

	if err := c.ShouldBindJSON(&movieUpdates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var movie models.Movie
	if err := config.DB.First(&movie, "id = ?", movieID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	// Update the movie
	if err := config.DB.Model(&movie).Updates(movieUpdates).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie updated successfully", "movie": movie})
}

func GetMostViewedMovie(c *gin.Context) {
	var movie models.Movie
	if err := config.DB.Order("view_count DESC").First(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch most viewed movie"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"movie": movie})
}

func GetMostViewedGenre(c *gin.Context) {
	type GenreView struct {
		Genre      string
		TotalViews int
	}

	var results []GenreView
	err := config.DB.Raw(`
        SELECT genre, SUM(view_count) AS total_views
        FROM (
            SELECT unnest(genres) AS genre, view_count
            FROM movies
        ) genre_views
        GROUP BY genre
        ORDER BY total_views DESC
        LIMIT 1
    `).Scan(&results).Error

	if err != nil || len(results) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch most viewed genre"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"most_viewed_genre": results[0]})
}

func ListMovies(c *gin.Context) {
	var movies []models.Movie
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	offset := (page - 1) * limit

	if err := config.DB.Offset(offset).Limit(limit).Find(&movies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"movies": movies})
}

func SearchMovies(c *gin.Context) {
	keyword := c.Query("keyword")

	var movies []models.Movie
	if err := config.DB.Where("title ILIKE ? OR description ILIKE ? OR artists::text ILIKE ? OR genres::text ILIKE ?",
		"%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").Find(&movies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search movies"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"movies": movies})
}

func TrackMovieView(c *gin.Context) {
	movieID := c.Param("id")
	var logData models.ViewLog

	if err := c.ShouldBindJSON(&logData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Fetch the movie
	var movie models.Movie
	if err := config.DB.First(&movie, "id = ?", movieID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	// Increment view count
	movie.ViewCount++
	if err := config.DB.Save(&movie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to track movie view"})
		return
	}

	// Log the view
	logData.MovieID = movieID
	if err := config.DB.Create(&logData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to log view"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "View tracked successfully"})
}

// ListVotedMovies lists all the movies a user has voted for.
func ListVotedMovies(c *gin.Context) {
	userID := c.GetString("user_id")

	var votes []models.Vote
	if err := config.DB.Where("user_id = ?", userID).Find(&votes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve votes"})
		return
	}

	var movieIDs []string
	for _, vote := range votes {
		movieIDs = append(movieIDs, vote.MovieID)
	}

	var movies []models.Movie
	if err := config.DB.Where("id IN ?", movieIDs).Find(&movies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve movies"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"movies": movies})
}
