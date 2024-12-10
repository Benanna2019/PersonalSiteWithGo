package helpers

import "os"

func GetSiteURL() string {
	env := os.Getenv("APP_ENV")
    if env == "development" {
        return "http://localhost:8080"
    }
    return "https://benpattonpersonalsite.fly.dev"
}
