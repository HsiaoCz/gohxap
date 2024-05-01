package etc

import "os"

func GetPort(key string) string {
	port := os.Getenv(key)
	if port == "" {
		return ":3001"
	}
	return port
}
