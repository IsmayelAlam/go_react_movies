package repository

import (
	"database/sql"
	"server/internal/models"
)

type DatabaseRepo interface {
	AddMovies() ([]*models.Movie, error)
	ConnectDb() *sql.DB
}
