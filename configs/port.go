package configs

import "os"

func GetPort() string {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8081"
	}

	return ":" + port
}
