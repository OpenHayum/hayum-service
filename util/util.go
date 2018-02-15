package util

// ConstructRoute constructs API endpoint
func ConstructEndpoint(basePath string, pathName string) string {
	return basePath + pathName
}
