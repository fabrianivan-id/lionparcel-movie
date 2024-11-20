# Movie Festival Backend

## Overview
This project is a backend service for a movie festival app. It allows admins to manage movies and users to vote for their favorite movies.

## Features
- Admin can add, update, and fetch most-viewed movies and genres.
- Users can view movies, vote, and unvote for movies.
- Authentication with JWT.
- Fully RESTful API.

## Tech Stack
- **Language**: Go
- **Framework**: Gin
- **Database**: PostgreSQL
- **Authentication**: JWT
- **ORM**: GORM

## Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/fabrianivan-id/movie-festival-backend.git

2. Install dependencies:
    ```go mod tidy
    Set up .env file:
    DATABASE_URL=your_database_url
    JWT_SECRET=your_secret
    SERVER_PORT=8080

3. Run the app:
    ```go run main.go```
