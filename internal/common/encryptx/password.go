package encryptx

import (
	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword  生成用户密码
func GeneratePassword(password string) (string, error) {
	pass := []byte(password)
	fromPassword, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(fromPassword), nil
}

// ComparePassword .
func ComparePassword(oldPassword, password string) error {
	oldPass := []byte(oldPassword)
	pass := []byte(password)
	return bcrypt.CompareHashAndPassword(oldPass, pass)
}
