package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tn-go-course/lynks/shortner/internal/models"
)

func (a *Api) GetShortURL(w http.ResponseWriter, r *http.Request) {
	u := &models.Url{}
	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		returnError(r.Context(), http.StatusBadRequest, err, w)
		return
	}

	u, err = a.app.CreateShortLink(r.Context(), u.Url)
	if err != nil {
		returnError(r.Context(), http.StatusBadRequest, err, w)
		return
	}

	returnData(r.Context(), u, w)
}

func (a *Api) GetURL(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	u, err := a.app.RestoreLink(r.Context(), id)
	if err != nil {
		returnError(r.Context(), http.StatusBadRequest, err, w)
		return
	}
	returnData(r.Context(), u, w)
}

func (a *Api) Redirect(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	destinationUrl, err := a.app.RestoreLink(r.Context(), id)
	if err != nil {
		returnError(r.Context(), http.StatusBadRequest, err, w)
		return
	}

	returnRedirect(w, r, destinationUrl.Url)
}
