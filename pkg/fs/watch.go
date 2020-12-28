package fs

import (
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
)

// Watch notifies for file changes for the given file or directory
func Watch(paths []string) error {

	// Initialize
	watcher, err := fsnotify.NewWatcher()
	defer watcher.Close()

	if err != nil {
		return err
	}

	// Set up a watcher for all paths
	for _, path := range paths {
		err = watcher.Add(path)
		if err != nil {
			return err
		}
	}

	select {
	case event, ok := <-watcher.Events:

		// Not 100% sure what a !ok event is
		// so debug for now
		if !ok {
			log.Debugf("Got a !ok event %v", event)
		}

	case err, ok := <-watcher.Errors:

		// Not 100% sure what an ok error is
		// so debug for now
		if ok {
			log.Debugf("Got an ok error %v", err)
		}

		// A very clear !ok plus error
		if !ok {
			return err
		}
	}

	return nil
}
