package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tn-go-course/lynks/memcache/internal/models"
)

func (a *Api) StoreURL(w http.ResponseWriter, r *http.Request) {
	el := &models.StorageElement{}
	err := json.NewDecoder(r.Body).Decode(&el)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	el, err = a.app.StoreUrl(r.Context(), el)
	if err != nil {
		returnError(r.Context(), http.StatusBadRequest, err, w)
		return
	}

	returnData(r.Context(), el, w)
}

func (a *Api) GetURL(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	u, err := a.app.GetUrl(r.Context(), id)
	if err != nil {
		returnError(r.Context(), http.StatusBadRequest, err, w)
		return
	}
	returnData(r.Context(), u, w)
}
