package employee

import (
	"database/sql"
)

type Repository struct {
	DB *sql.DB
}

// GetProfile ambil profil lengkap karyawan + nama divisi + nama cabang
func (r *Repository) GetProfile(employeeID int) (*ProfileResponse, error) {
	var p ProfileResponse

	err := r.DB.QueryRow(`
		SELECT 
			e.id, e.name, e.username, e.role, e.tipe,
			COALESCE(d.name, '-') AS division_name,
			COALESCE(b.name, '-') AS branch_name
		FROM employees e
		LEFT JOIN divisions d ON d.id = e.division_id
		LEFT JOIN branches b  ON b.id = e.branch_id
		WHERE e.id = $1
	`, employeeID).Scan(
		&p.ID, &p.Name, &p.Username, &p.Role, &p.Tipe,
		&p.DivisionName, &p.BranchName,
	)

	if err != nil {
		return nil, err
	}
	return &p, nil
}

// GetPasswordHash ambil hash password karyawan untuk verifikasi
func (r *Repository) GetPasswordHash(employeeID int) (string, error) {
	var hash string
	err := r.DB.QueryRow(`
		SELECT password FROM employees WHERE id = $1
	`, employeeID).Scan(&hash)
	return hash, err
}

// UpdatePassword update password karyawan dengan hash baru
func (r *Repository) UpdatePassword(employeeID int, newHash string) error {
	_, err := r.DB.Exec(`
	UPDATE employees SET password = $1 WHERE id = $2
	`, newHash, employeeID)
	return err
}
