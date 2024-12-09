package helpers

import "os"

func GetSiteURL() string {
	env := os.Getenv("APP_ENV")
    if env == "development" {
        return "http://localhost:3000"
    }
    return "https://benapatton.com"
}
