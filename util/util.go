package util

import (
	"path"
)

// ConstructRoute constructs API endpoint
func ConstructEndpoint(basePath string, pathName string) string {
	return path.Join(basePath, pathName)
}
