/*
Package user provides the user service routes
Testing the user service handlers
*/

package user

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/ykkalexx/e-com-backend/types"
)

func TestUserServiceHandlers(t *testing.T) {
	userStore := &mockUserStore{}
	handler := NewHandler(userStore)
	t.Run(" ail if the user payload is invalid", func(t *testing.T) {
		payload := types.NewUser{
			FirstName: "test",
			LastName: "test",
			Email: "",
			Password: "test",
		}
		marshalled, _ := json.Marshal(payload)

		req, err  := http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(marshalled))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		router := mux.NewRouter()
		router.HandleFunc("/register", handler.handleRegister)

		router.ServeHTTP(rr, req)

		if rr.Code != http.StatusBadRequest {
			t.Errorf("expected status %d; got %d", http.StatusBadRequest, rr.Code)
		}
	})
}

type mockUserStore struct {}

func (m *mockUserStore) FetchUserByEmail(email string) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) FetchUserByID(id int) (*types.User, error) {
	return nil, nil
}

func (m *mockUserStore) CreateUser(u types.User) error {
	return nil
}