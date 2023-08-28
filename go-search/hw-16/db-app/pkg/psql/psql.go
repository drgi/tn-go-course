package psql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	database "github.com/tn-go-course/go-search/hw-16/db-app/pkg"
)

type DB struct {
	p *pgxpool.Pool
}

func New(url string) (*DB, error) {
	ctx := context.Background()
	pool, err := pgxpool.New(ctx, url)
	if err != nil {
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		return nil, err
	}

	return &DB{pool}, nil
}

func (db *DB) Films(ctx context.Context, f *database.FilterFilm) ([]*database.Film, error) {
	sql := "SELECT * FROM films"

	if f.StudioID > 0 {
		sql += fmt.Sprintf(" WHERE studio_id = %d", f.StudioID)
	}

	rows, err := db.p.Query(ctx, sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	films := make([]*database.Film, 0)

	for rows.Next() {
		film := &database.Film{}
		err := rows.Scan(
			&film.ID,
			&film.Name,
			&film.RealiseYear,
			&film.Profit,
			&film.Rating,
			&film.StudioID,
		)
		if err != nil {
			return nil, err
		}
		films = append(films, film)
	}

	return films, nil
}

func (db *DB) CreateFilms(ctx context.Context, films []*database.Film) error {
	batch := &pgx.Batch{}
	for _, film := range films {
		batch.Queue(
			"INSERT INTO films (name, realise_year, profit, rating, studio_id) VALUES ($1, $2, $3, $4, $5)",
			film.Name, film.RealiseYear, film.Profit, film.Rating, film.StudioID,
		)
	}

	return db.p.SendBatch(ctx, batch).Close()
}

func (db *DB) DeleteFilm(ctx context.Context, id uint) (err error) {
	tx, err := db.p.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, "DELETE FROM films_actors WHERE film_id = $1", id)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "DELETE FROM films_directors WHERE film_id = $1", id)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "DELETE FROM films WHERE id = $1", id)
	if err != nil {
		return err
	}
	return
}

func (db *DB) UpdateFilm(ctx context.Context, film *database.Film) (err error) {
	_, err = db.p.Exec(ctx,
		"UPDATE films SET name = $1, realise_year = $2, profit = $3, rating = $4, studio_id = $5 WHERE id = $6",
		film.Name, film.RealiseYear, film.Profit, film.Rating, film.StudioID, film.ID,
	)
	return
}
