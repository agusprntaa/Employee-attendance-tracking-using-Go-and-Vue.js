package auth

import (
	"database/sql"
	"time"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) FindUser(username string) (int, string, string, error) {
	var id int
	var passsword, role string

	err := r.DB.QueryRow(`
		SELECT id, password, role
		FROM employees
		WHERE username = $1
		`, username).Scan(&id, &passsword, &role)

	return id, passsword, role, err
}

func (r *Repository) SaveRefreshToken(userID int, token string, exp time.Time) error {
	_, err := r.DB.Exec(`
		INSERT INTO refresh_tokens (employee_id, token, expires_at)
		VALUES ($1, $2, $3)
		`, userID, token, exp)

	return err
}

func (r *Repository) ValidateRefreshToken(token string) (int, error) {
	var userID int

	err := r.DB.QueryRow(`
		SELECT employee_id
		FROM refresh_tokens
		WHERE token = $1 AND expires_at > NOW()
		`, token).Scan(&userID)

	return userID, err
}

func (r *Repository) DeleteRefreshToken(token string) {
	r.DB.Exec(`DELETE FROM refresh_tokens WHERE token = $1`, token)
}

func (r *Repository) CreateUser(username, password, role, tipe string) error {
	_, err := r.DB.Exec(`
		INSERT INTO employees (username, password, role, tipe)
		VALUES ($1, $2, $3, $4)
		`, username, password, role, tipe)

	return err
}
