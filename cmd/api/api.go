package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/navyn13/PingMySite/services/url"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Home Page")
	}).Methods("GET")

	urlHandler := url.NewHandler()
	urlHandler.RegisterRoutes(router)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
