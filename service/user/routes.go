package user

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

// RegisterRoutes registers the routes for the user service
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleLogin).Methods("POST")
}

// handleLogin handles the login request
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

// handleRegister handles the register request
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {

}