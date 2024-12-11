package helpers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv" // Add this import
)

func GetSiteURL() string {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}
	env := os.Getenv("APP_ENV")
    if env == "development" {
        return "http://localhost:8080"
    }
    return "https://benpattonpersonalsite.fly.dev"
}

func GetCacheControl() string {
    if os.Getenv("APP_ENV") == "production" {
        return "max-age=3600" // 1 hour
    }
    return "no-cache, no-store, must-revalidate, max-age=0, s-maxage=0"
}