package helpers

import (
	"log"

	"github.com/fsnotify/fsnotify"
)

func WatchContent() error {
    watcher, err := fsnotify.NewWatcher()
    if err != nil {
        return err
    }
    defer watcher.Close()

    go func() {
        for {
            select {
            case event, ok := <-watcher.Events:
                if !ok {
                    return
                }
                if event.Has(fsnotify.Write) || event.Has(fsnotify.Create) || event.Has(fsnotify.Remove) {
                    log.Println("Content changed, regenerating RSS...")
                    if err := GenerateRSSFeed(); err != nil {
                        log.Printf("Error generating RSS: %v\n", err)
                    }
                }
            case err, ok := <-watcher.Errors:
                if !ok {
                    return
                }
                log.Println("Error:", err)
            }
        }
    }()

    err = watcher.Add("./content")
    if err != nil {
        return err
    }

    // Keep running
    select {}
}