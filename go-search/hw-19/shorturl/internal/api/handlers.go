package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Data struct {
	Url      string `json:"url"`
	ShortURL string `json:"shortUrl"`
}

func (a *Api) GetShortURL(w http.ResponseWriter, r *http.Request) {
	data := &Data{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data.ShortURL, err = a.app.CreateShortLink(r.Context(), data.Url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (a *Api) GetURL(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	url, err := a.app.RestoreLink(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res := &Data{
		Url:      url,
		ShortURL: id,
	}
	err = json.NewEncoder(w).Encode(&res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (a *Api) Redirect(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	url, err := a.app.RestoreLink(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, url, http.StatusSeeOther)
}
