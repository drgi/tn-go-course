package webapp

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Sercher interface {
	List() (interface{}, error)
	Documents() (interface{}, error)
}

type WebServer struct {
	router *mux.Router
	index  Sercher
}

func New(s Sercher) *WebServer {
	router := mux.NewRouter()
	return &WebServer{
		router: router,
		index:  s,
	}
}

func (s *WebServer) RegisterHandlers() {
	s.router.HandleFunc("/index", s.Index).Methods(http.MethodGet)
	s.router.HandleFunc("/docs", s.Documents).Methods(http.MethodGet)
}

func (s *WebServer) Listen(addr string) error {
	return http.ListenAndServe(addr, s.router)
}
