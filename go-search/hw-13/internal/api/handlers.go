package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tn-go-course/go-search/hw-13/pkg/crawler"
)

const (
	queryKey = "query"
)

func (s *Api) Search(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get(queryKey)
	if query == "" {
		returnError(r.Context(), http.StatusBadRequest, errors.New("query is required"), w)
		return
	}
	docs, err := s.app.Search(r.Context(), query)
	if err != nil {
		returnError(r.Context(), http.StatusBadRequest, err, w)
		return
	}
	returnData(r.Context(), docs, w)
}

func (s *Api) CreateDocument(w http.ResponseWriter, r *http.Request) {
	params := crawler.Document{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		returnError(r.Context(), http.StatusBadRequest, err, w)
		return
	}
	ctx := r.Context()
	id, err := s.app.CreateDocument(ctx, params)
	if err != nil {
		returnError(r.Context(), http.StatusBadRequest, err, w)
		return
	}
	returnData(ctx, id, w)
}

func (s *Api) UpdateDocument(w http.ResponseWriter, r *http.Request) {
	params := crawler.Document{}
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		returnError(r.Context(), http.StatusBadRequest, err, w)
		return
	}
	ctx := r.Context()
	err = s.app.UpdateDocument(ctx, params)
	if err != nil {
		returnError(r.Context(), http.StatusBadRequest, err, w)
		return
	}
	returnData(ctx, true, w)
}

func (s *Api) DeleteDocument(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	id, err := strconv.Atoi(param["id"])
	if err != nil {
		returnError(r.Context(), http.StatusBadRequest, err, w)
		return
	}
	ctx := r.Context()
	err = s.app.DeleteDocument(ctx, id)
	if err != nil {
		returnError(r.Context(), http.StatusBadRequest, err, w)
		return
	}
	returnData(ctx, true, w)
}
