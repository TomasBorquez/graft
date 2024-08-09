package graft

import (
	"github.com/TomasBorquez/logger"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func AddWatcher(opts *HRConfig, cb TaskCb) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer func(watcher *fsnotify.Watcher) {
		err := watcher.Close()
		if err != nil {
			logger.Error("[Graft]: Error meanwhile closing watcher: %v", err)
		}
	}(watcher)
	
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				
				if event.Has(fsnotify.Write) {
					logger.Info(`[Graft]: Modified file "%s", running script again...`, event.Name)
					cb(&TaskConfig{})
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	
	err = addRecursiveWatch(watcher, opts)
	if err != nil {
		log.Fatal(err)
	}
	
	// Waits forever
	<-make(chan struct{})
}

func addRecursiveWatch(watcher *fsnotify.Watcher, opts *HRConfig) error {
	return filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		path = convertPath(path)
		if err != nil {
			return err
		}
		
		if shouldIgnore(path, opts) {
			if info.IsDir() {
				return filepath.SkipDir
			} else {
				return nil
			}
		}
		
		if info.IsDir() {
			return nil
		}
		return watcher.Add(path)
	})
}

func convertPath(path string) string {
	return strings.ReplaceAll(path, "\\", "/")
}

func shouldIgnore(path string, opts *HRConfig) bool {
	// Check if the path is in ExcludeDir
	if isDir(path) {
		for _, dir := range opts.ExcludeDir {
			if path == dir {
				return true
			}
		}
		
		if len(opts.IncludeDir) != 0 {
			for _, dir := range opts.IncludeDir {
				if path == dir {
					return false
				}
			}
			return true
		}
		
		return false
	}
	
	ext := filepath.Ext(path)
	for _, excludeExt := range opts.ExcludeExten {
		if ext == excludeExt {
			return true
		}
	}
	
	if len(opts.IncludeExten) != 0 {
		for _, includeExt := range opts.IncludeExten {
			if ext == includeExt {
				return false
			}
		}
		return true
	}
	
	return false
}

func isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}
