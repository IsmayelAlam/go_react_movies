package dbrepo

import (
	"context"
	"database/sql"
	"server/internal/models"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 5 // seconds

func (m *PostgresDBRepo) ConnectDb() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) AddMovies() ([]*models.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	qurey := `SELECT
	id, title, description, release_date, runtime, mpaa_rating, coalesce(image, ''), created_at, updated_at
	FROM movies
	ORDER BY title
	`

	rows, err := m.DB.QueryContext(ctx, qurey)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*models.Movie

	for rows.Next() {
		var m models.Movie
		err := rows.Scan(
			&m.ID,
			&m.Title,
			&m.Description,
			&m.ReleaseDate,
			&m.RunTime,
			&m.MpaaRating,
			&m.Image,
			&m.CreatedAt,
			&m.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		movies = append(movies, &m)
	}

	return movies, nil
}
