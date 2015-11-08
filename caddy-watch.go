package main

import (
	"fmt"
	"gopkg.in/fsnotify.v1"
	"log"
)

const watch_path = "./app"

func main() {

	fmt.Println("connected")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				if event.Op&fsnotify.Write == fsnotify.Write {
					fmt.Println(event.Name)
				}
			case err := <-watcher.Errors:
				fmt.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(watch_path)

	if err != nil {
		log.Fatal(err)
	}
	<-done
}
