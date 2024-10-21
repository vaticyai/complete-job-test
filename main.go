package main

import (
	"os"

	"github.com/vaticyai/syncpocalypse/backup"
	"github.com/vaticyai/syncpocalypse/logger"
)

const (
	source      string = "./_source/crawl.md"      // Source file to copy
	destination string = "./_destination/crawl.md" // The destination to copy to
)

func main() {
	err := backup.Run(source, destination)
	if err != nil {
		logger.Logger.Errorf("Error backing file '%s' to '%s': %v", source, destination, err)
		os.Exit(1)
	} else {
		logger.Logger.Debug("File copied successfully!")
	}
}
