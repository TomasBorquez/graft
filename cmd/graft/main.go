package main

import (
	"github.com/TomasBorquez/graft/internal/store"
	"github.com/TomasBorquez/logger"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printHelp()
		return
	}
	
	action := args[0]
	
	switch action {
	case "start":
		store.ExecuteHandlers("start")
	case "build":
		store.ExecuteHandlers("build")
	case "test":
		store.ExecuteHandlers("test")
	case "help":
		printHelp()
	default:
		logger.Error("Command %s does not exist, run `graft help` to see the list of commands", action)
	}
}

func printHelp() {
	logger.Warning("TODO")
}