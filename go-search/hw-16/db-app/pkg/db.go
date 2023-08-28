package database

import "context"

type Interface interface {
	Films(context.Context, *FilterFilm) ([]*Film, error)
	CreateFilms(context.Context, []*Film) error
	DeleteFilm(context.Context, uint) error
	UpdateFilm(context.Context, *Film) error
}

type FilterFilm struct {
	StudioID uint
}

type Film struct {
	ID          uint
	Name        string
	RealiseYear uint
	Profit      int
	Rating      string
	StudioID    uint
}
