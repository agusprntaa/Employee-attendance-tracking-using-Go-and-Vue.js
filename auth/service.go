package auth

import (
	"absensi_karyawan/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repo *Repository
}

// User struct untuk dikirim ke handler (dan ke frontend)
type User struct {
	ID           int
	Name         string
	Username     string
	Role         string
	EmployeeType string
}

func (s *Service) Login(username, password string) (string, string, *User, error) {
	user, hashed, err := s.Repo.FindUser(username)
	if err != nil {
		return "", "", nil, errors.New("user not found")
	}

	// bandingkan password
	if bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)) != nil {
		return "", "", nil, errors.New("wrong password")
	}

	// generate tokens
	access, _ := utils.GenerateAccessToken(user.ID, user.Role)
	refresh, exp, _ := utils.GenerateRefreshToken(user.ID)

	// simpan refresh token ke DB
	s.Repo.SaveRefreshToken(user.ID, refresh, exp)

	return access, refresh, user, nil
}

func (s *Service) Refresh(oldToken string) (string, error) {
	userID, role, err := s.Repo.ValidateRefreshToken(oldToken)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	newAccess, _ := utils.GenerateAccessToken(userID, role)

	return newAccess, nil
}

func (s *Service) Logout(refreshToken string) {
	s.Repo.DeleteRefreshToken(refreshToken)
}

func (s *Service) CreateUser(username, password, name, role, tipe string) error {
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	return s.Repo.CreateUser(username, hashed, name, role, tipe)
}
