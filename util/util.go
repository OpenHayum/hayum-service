package util

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"log"
	"math/rand"
	"net/http"
	"path"
	"strconv"
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

// GetRandID generate random string
func GetRandID() string {
	return uuid.New().String()
}

func StrToInt64(str string) (int64, error) {
	i, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		return 0, err
	}

	return i, nil
}
