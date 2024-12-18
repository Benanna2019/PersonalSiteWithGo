package helpers

import (
	"bytes"
	"encoding/base64"
	"log"
	"os"

	// "path/filepath"
	"sort"
	"time"

	"github.com/andybalholm/brotli" // Add this line
	"github.com/gorilla/feeds"
)

func getHostname() string {
    return GetSiteURL(); 
}

func GenerateRSSFeed() error {
    hostname := getHostname()

    // Create new feed
    feed := &feeds.Feed{
        Title:       "Enhance Blog Template",
        Link:        &feeds.Link{Href: hostname},
        Description: "My blog description.",
        Copyright:   "All rights reserved " + time.Now().Format("2006") + ", My Company",
        Created:     time.Now(),
        Author: &feeds.Author{
            Name: "My Company",
        },
    }

    // Get all posts
    posts, err := GetPostMetadata()
    if err != nil {
        return err
    }

    // Sort posts by date (newest first)
    sort.Slice(posts, func(i, j int) bool {
        return posts[i].Frontmatter.Published > posts[j].Frontmatter.Published
    })

    // Add items to feed
    for _, post := range posts {
        // Get full post content
        postContent, err := GetPostById(post.Frontmatter.ID)
        if err != nil {
            continue
        }

        pubDate, err := time.Parse("January 2, 2006", post.Frontmatter.Published)
        if err != nil {
            log.Printf("Error parsing date for post %s: %v", post.Frontmatter.ID, err)
            continue
        }

        item := &feeds.Item{
            Title:       post.Frontmatter.Title,
            Link:        &feeds.Link{Href: hostname + "/posts/" + post.Frontmatter.Slug},
            Description: post.Frontmatter.Description,
            Content:     postContent.Html,
            Author:      &feeds.Author{Name: post.Frontmatter.Author},
            Created:     pubDate,
            Id:         hostname + "/posts/" + post.Frontmatter.Slug, // Add GUID
        }

        feed.Items = append(feed.Items, item)
    }

    // Generate RSS 2.0 feed
    rss, err := feed.ToRss()
    if err != nil {
        return err
    }

    // Save uncompressed XML
    err = os.WriteFile("web/static/rss.xml", []byte(rss), 0644)
    if err != nil {
        return err
    }

    // Create Brotli compressed version
    var compressed bytes.Buffer
    bw := brotli.NewWriter(&compressed)
    if _, err := bw.Write([]byte(rss)); err != nil {
        return err
    }
    if err := bw.Close(); err != nil {
        return err
    }

    // Save compressed version
    compressedData := base64.StdEncoding.EncodeToString(compressed.Bytes())
    return os.WriteFile("web/static/rss.br", []byte(compressedData), 0644)
}

func parseDate(dateStr string) time.Time {
    // Add your date parsing logic here based on your date format
    t, _ := time.Parse("January 2, 2006", dateStr)
    return t
}