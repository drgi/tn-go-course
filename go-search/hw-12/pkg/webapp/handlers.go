package webapp

import (
	"encoding/json"
	"net/http"
)

func (s *WebServer) Index(w http.ResponseWriter, r *http.Request) {
	data, err := s.index.List()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	response, err := json.MarshalIndent(data, "", "  ")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func (s *WebServer) Documents(w http.ResponseWriter, r *http.Request) {
	data, err := s.index.Documents()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	response, err := json.MarshalIndent(data, "", "  ")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
