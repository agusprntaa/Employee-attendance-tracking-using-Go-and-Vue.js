package auth

import (
	"absensi_karyawan/utils"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service struct {
	Repo *Repository
}

func (s *Service) Login(username, password string) (string, string, error) {
	id, hashed, role, err := s.Repo.FindUser(username)
	if err != nil {
		return "", "", errors.New("user not found")
	}

	// compare password
	if bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)) != nil {
		return "", "", errors.New("wrong password")
	}

	// generate tokens
	access, _ := utils.GenerateAccessToken(id, role)
	refresh, exp, _ := utils.GenerateRefreshToken(id)

	// simpan refresh token
	s.Repo.SaveRefreshToken(id, refresh, exp)

	return access, refresh, nil
}

func (s *Service) Refresh(oldToken string) (string, error) {
	userID, err := s.Repo.ValidateRefreshToken(oldToken)
	if err != nil {
		return "", errors.New("invalid refresh token")
	}

	// ⚠️ NOTE: ini masih hardcode
	newAccess, _ := utils.GenerateAccessToken(userID, "karyawan")

	return newAccess, nil
}

func (s *Service) Logout(refreshToken string) {
	s.Repo.DeleteRefreshToken(refreshToken)
}

func (s *Service) CreateUser(username, password, role, tipe string) error {
	hashed, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	return s.Repo.CreateUser(username, hashed, role, tipe)
}
