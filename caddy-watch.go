package main

import (
	"fmt"
	"gopkg.in/fsnotify.v1"
	"log"
	"os"
	"path/filepath"
)

// Let's do this
// take in paths as CL Args
// walk the paths
func main() {

	fmt.Println("caddy-hot-watcher connected")

	done := make(chan bool)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

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

	root := "./"
	err = filepath.Walk(root, watchThis(*watcher, root))
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func watchThis(watcher fsnotify.Watcher, root string) filepath.WalkFunc {
	return func(path string, info os.FileInfo, err error) error {

		// skip these dirs
		if path == "node_modules" ||
			path == "jspm_packages" ||
			path == ".git" {
			return filepath.SkipDir
		}
		// recurse unless self
		if info.IsDir() && path != root {
			return filepath.Walk(path, watchThis(watcher, path))
		}

		return watcher.Add(path)
	}
}
