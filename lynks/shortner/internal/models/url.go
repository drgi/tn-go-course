package models

type Url struct {
	// PG primary key
	ID uint `json:"id"`
	// Destination url
	Url string `json:"url"`
	// Short string ID
	StringID string `json:"stringId"`
}
