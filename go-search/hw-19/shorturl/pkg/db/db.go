package db

type DB interface {
	StoreUrl(*Url) error
	RestoreUrl(string) (*Url, error)
}

type Url struct {
	Url string
	Id  string
}
