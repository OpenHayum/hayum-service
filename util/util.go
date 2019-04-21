package util

import (
	"log"
	"math/rand"
	"net/http"
	"path"

	"golang.org/x/crypto/bcrypt"
)

// ConstructEndpoint constructs API endpoint
func ConstructEndpoint(basePath string, pathName string) string {
	return path.Join(basePath, pathName)
}

// GenerateOTP generates OTP which will be used for SMS OTP verification
func GenerateOTP() int32 {
	return rand.Int31n(10000)
}

// EncryptPassword encryptes password
func EncryptPassword(password string) (string, error) {
	if password == "" {
		return "", nil
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

// CompareHashAndPassword compares hashed password and actual password
func CompareHashAndPassword(hashedPassword, password string) error {
	byteHashedPassword := []byte(hashedPassword)
	bytePassword := []byte(password)

	return bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword)
}

// GetCookieValue extracts cookie value by name
func GetCookieValue(r *http.Request, name string) string {
	cookie, err := r.Cookie(name)
	if err != nil {
		log.Println("getCookieValue: Unable to read cookie with name: ", name, err)
		return ""
	}
	return cookie.Value
}
