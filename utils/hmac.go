package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"os"
)

// GenerateQRToken membuat token HMAC untuk QR code harian
// Format payload: "branchID|tanggal" contoh: "1|2026-04-27"
func GenerateQRToken(branchID int, date string) string {
	payload := fmt.Sprintf("%d|%s", branchID, date)
	secret := []byte(os.Getenv("HMAC_SECRET"))

	mac := hmac.New(sha256.New, secret)
	mac.Write([]byte(payload))

	return base64.URLEncoding.EncodeToString(mac.Sum(nil))
}

// ValidateQRToken memvalidasi token QR yang dikirim karyawan saat check-in
func ValidateQRToken(token string, branchID int, date string) bool {
	expected := GenerateQRToken(branchID, date)
	// pakai hmac.Equal untuk mencegah timing attack
	return hmac.Equal([]byte(token), []byte(expected))
}
