package auth

import (
	"database/sql"
	"time"
)

type Repository struct {
	DB *sql.DB
}

func (r *Repository) FindUser(username string) (*User, string, error) {
	var user User
	var hashed string

	err := r.DB.QueryRow(`
		SELECT id, name, username, role, tipe, password
		FROM employees
		WHERE username = $1
	`, username).Scan(
		&user.ID,
		&user.Name,
		&user.Username,
		&user.Role,
		&user.EmployeeType,
		&hashed,
	)

	if err != nil {
		return nil, "", err
	}

	return &user, hashed, nil
}

func (r *Repository) SaveRefreshToken(userID int, token string, exp time.Time) error {
	_, err := r.DB.Exec(`
		INSERT INTO refresh_tokens (employee_id, token, expires_at)
		VALUES ($1, $2, $3)
	`, userID, token, exp)

	return err
}

func (r *Repository) ValidateRefreshToken(token string) (int, string, error) {
	var userID int
	var role string

	err := r.DB.QueryRow(`
		SELECT rt.employee_id, e.role
		FROM refresh_tokens rt
		JOIN employees e ON e.id = rt.employee_id
		WHERE rt.token = $1 AND rt.expires_at > NOW()
	`, token).Scan(&userID, &role)

	return userID, role, err
}

func (r *Repository) DeleteRefreshToken(token string) {
	r.DB.Exec(`DELETE FROM refresh_tokens WHERE token = $1`, token)
}

func (r *Repository) CreateUser(username, password, name, role, tipe string) error {
	_, err := r.DB.Exec(`
		INSERT INTO employees (username, password, name, role, tipe)
		VALUES ($1, $2, $3, $4, $5)
	`, username, password, name, role, tipe)

	return err
}
