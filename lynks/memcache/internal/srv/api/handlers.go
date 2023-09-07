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

	err = a.cache.SetString(r.Context(), el.Key, el.Value)
	if err != nil {
		returnError(r.Context(), http.StatusBadRequest, err, w)
		return
	}

	returnData(r.Context(), el, w)
}

func (a *Api) GetURL(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	u, err := a.cache.GetString(r.Context(), id)
	if err != nil {
		returnError(r.Context(), http.StatusBadRequest, err, w)
		return
	}
	returnData(r.Context(), u, w)
}
