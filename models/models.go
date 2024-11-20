package models

import "time"

type Movie struct {
	ID          string   `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Title       string   `gorm:"type:varchar(255);not null" json:"title"`
	Description string   `gorm:"type:text" json:"description"`
	Duration    int      `json:"duration"`
	Artists     []string `gorm:"type:jsonb" json:"artists"`
	Genres      []string `gorm:"type:jsonb" json:"genres"`
	WatchURL    string   `gorm:"type:varchar(255);not null" json:"watch_url"`
	ViewCount   int      `gorm:"default:0" json:"view_count"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type User struct {
	ID           string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	Username     string `gorm:"type:varchar(50);unique;not null" json:"username"`
	Email        string `gorm:"type:varchar(100);unique;not null" json:"email"`
	PasswordHash string `gorm:"type:varchar(255);not null" json:"-"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Vote struct {
	ID        string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    string `gorm:"type:uuid;not null" json:"user_id"`
	MovieID   string `gorm:"type:uuid;not null" json:"movie_id"`
	CreatedAt time.Time
}

type ViewLog struct {
	ID              string `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	MovieID         string `gorm:"type:uuid;not null" json:"movie_id"`
	UserID          string `gorm:"type:uuid" json:"user_id"`
	DurationWatched int    `json:"duration_watched"`
	CreatedAt       time.Time
}
