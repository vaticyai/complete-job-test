package main

import (
	"fmt"
	"os"

	"github.com/vaticyai/syncpocalypse/backup"
	"github.com/vaticyai/syncpocalypse/logger"
	"github.com/vaticyai/syncpocalypse/metrics"
)

const (
	source      string = "./_source/crawl.md"         // Source file to copy
	destination string = "./_destination/crawl-%d.md" // The destination to copy to
)

func main() {
	go metrics.Run()

	for i := 1; i < 5_000; i++ {
		err := backup.Run(source, fmt.Sprintf(destination, i))
		if err != nil {
			logger.Logger.Errorf("Error backing file '%s' to '%s': %v", source, destination, err)
			metrics.TotalBackups.WithLabelValues("error").Inc()
			os.Exit(1)
		} else {
			logger.Logger.Debug("File copied successfully!")
			metrics.TotalBackups.WithLabelValues("success").Inc()
		}
	}
}
