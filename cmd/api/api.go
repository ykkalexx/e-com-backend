package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ykkalexx/e-com-backend/service/user"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

// NewAPIServer creates a new instance of APIServer
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

// Run starts the API server
func (s *APIServer) Run() error {
	// Create a new router
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// Create a new user store
	// dependency injection of the database connection 
	userStore := user.NewStore(s.db)

	// Register the routes for the user service
	userService := user.NewHandler(userStore)
	userService.RegisterRoutes(subrouter)

	// Start the server
	log.Println("Starting server on", s.addr)

	// Listen and serve
	return http.ListenAndServe(s.addr, router)
}
