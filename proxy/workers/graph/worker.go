package graph

import (
	"log"
	"os"
	"time"
)

const (
	appPathToGraphFile = "/app/static/tasks/graph.md"
)

func GraphWorker() {
	t := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-t.C:
			err := os.WriteFile(appPathToGraphFile, []byte(generateMarkdownText()), 0644)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
