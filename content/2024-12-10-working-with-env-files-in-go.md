---
"_id": "cm4j8ip0800000cmhh14vel4r"
title: "Working with .env files in Go "
published: "December 10, 2024"
slug: "working-with-env-files-in-go"
description: "Using .env files in Go"
categories: ["Go", "Enhance", "Extism", "Fly.io"]
author: "Ben Patton"
authorImage: "https://cdn.sanity.io/images/nfspldzq/production/3ceeda54221c7c0614ecc51f955c7be39a1da34e-512x512.jpg"
type: "Blog"
---

_This is part of a format I am working towards for short form content/snippets to reference later. I got this idea from R. Mark Volkmann on an episode of the [hypermedia-tv](https://www.youtube.com/watch?v=bs4UZR6c4Qc&t) podcast._

This blog is a migration to Go so there are a lot of things I am learning as I am converting this blog to Go.

One of these was in regards to setting Cache Control headers based off of the environment. Initially I was under the impression that the environment is available similar to a node.js environment. Below is the way I had it setup but I had to do a couple things

```go

import "os"

func GetSiteURL() string {
	env := os.Getenv("APP_ENV")
    if env == "development" {
        return "http://localhost:8080"
    }
    return "https://benpattonpersonalsite.fly.dev"
}

```

First I had to add a package to load the env which was `github.com/joho/godotenv`. Then I added the `godotenv.Load()` to the top of the code block and added some error handling/logging just in case. I also had to actually have a .env file with the APP_ENV variable set to either development or production. So locally I would have it set to development and on [fly.io](https://fly.io), where I am deploying to, it would be set to production.

```go
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
```

As for setting the env variable, locally it just goes in the .env and reads APP_ENV=development. For setting an app secret on fly.io, I use the command line and run

```bash
fly secrets set APP_ENV=production
```

Since fly is connected to this project when I set it up with the command line, I don't have to specify anything else other than the name of the variable.
