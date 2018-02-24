package util

import (
	"math/rand"
	"path"
)

// ConstructEndpoint constructs API endpoint
func ConstructEndpoint(basePath string, pathName string) string {
	return path.Join(basePath, pathName)
}

// GenerateOTP generates OTP which will be used for SMS OTP verification
func GenerateOTP() int {
	return rand.Intn(10000)
}
