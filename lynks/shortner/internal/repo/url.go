package repo

import (
	"context"

	"github.com/tn-go-course/lynks/shortner/internal/models"
)

const TableNameUrls = "urls"

func (r *Repo) CreateUrl(ctx context.Context, u *models.Url) (uint, error) {
	var id uint
	err := r.DB.
		QueryRow(ctx, "INSERT INTO urls (url, string_id) VALUES ($1, $2) RETURNING id", u.Url, u.StringID).
		Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *Repo) GetUrlByStringId(ctx context.Context, id string) (*models.Url, error) {
	u := &models.Url{}
	row := r.DB.QueryRow(ctx, "SELECT * FROM urls WHERE string_id = $1", id)
	err := row.Scan(&u.ID, &u.Url, &u.StringID)
	if err != nil {
		return nil, err
	}
	return u, nil
}
