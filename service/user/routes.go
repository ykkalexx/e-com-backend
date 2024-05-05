package user

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ykkalexx/e-com-backend/service/auth"
	"github.com/ykkalexx/e-com-backend/types"
	"github.com/ykkalexx/e-com-backend/utils"
)

type Handler struct {
	store types.UserStore
}

func NewHandler(store types.UserStore) *Handler {
	return &Handler{store: store}
}

// RegisterRoutes registers the routes for the user service
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

// handleLogin handles the login request
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

// handleRegister handles the register request
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
    // json payload should contain the user details
    var payload types.NewUser
    if err := utils.ParseJSON(r, &payload); err != nil {
        utils.WriteError(w, http.StatusBadRequest, err)
        return
    }
    // if the user already exists return an error
    user, err := h.store.FetchUserByEmail(payload.Email)
    if err != nil && err.Error() != "user not found" {
        utils.WriteError(w, http.StatusInternalServerError, err)
        return
    }
    if user != nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exist", payload.Email))
        return
    }

    hashedPassword, err := auth.HashedPassword(payload.Password);
    if err != nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error hashing password"))
        return
    }

    // if it doesn't exist, create a new user 
    if err := h.store.CreateUser(types.User{
        FirstName: payload.FirstName,
        LastName:  payload.LastName,
        Email:     payload.Email,
        Password:  hashedPassword,
    }); err != nil {
        utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("error creating %s user", payload.Email))
        return
    }

    utils.WriteJSON(w, http.StatusCreated, nil)
}