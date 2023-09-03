package repo

import (
	"github.com/rs/zerolog"
	"github.com/tn-go-course/lynks/shortner/pkg/postgres"
)

type Repo struct {
	*postgres.DB
	logger *zerolog.Logger
}

func New(db *postgres.DB, logger *zerolog.Logger) *Repo {
	r := &Repo{}
	r.DB = db
	r.logger = logger
	return r
}
