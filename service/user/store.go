package user

import (
	"database/sql"
	"fmt"

	"github.com/ykkalexx/e-com-backend/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

// FetchUserByEmail fetches a user by email from the database
func (s *Store) FetchUserByEmail(email string) (*types.User, error) {
	// query the database for the user with the given email
	rows, err := s.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return nil, err
	}

	// scan the result into a User struct
	u := new(types.User)
	for rows.Next() {
		u, err = ScanRowIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}

	// if the user is not found, return an error
	if u.ID == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return u, nil
}

// FetchUserByID fetches a user by ID from the database
func (s *Store) FetchUserByID(id int) (*types.User, error) {
	return nil, nil
}

// CreateUser creates a new user in the database
func (s *Store) CreateUser(u types.User) error {
	return nil
}

// ScanRowIntoUser scans the current row of the result set into a User struct
func ScanRowIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)

	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}