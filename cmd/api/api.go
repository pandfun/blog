package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/pandfun/blog/service/post"
	"github.com/pandfun/blog/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{addr: addr, db: db}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Registering user routes
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	// Registering post routes
	postStore := post.NewStore(s.db)
	postHandler := post.NewHandler(postStore)
	postHandler.RegisterRoutes(subrouter)

	log.Println("Server: Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
