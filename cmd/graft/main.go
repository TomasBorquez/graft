package main

import (
	"github.com/TomasBorquez/graft/internal/store"
	"github.com/TomasBorquez/logger"
	"os"
)

func main() {
	args := os.Args[1:]
	action := args[0]
	
	switch action {
	case "start":
		graft.ExecuteHandlers("start")
	case "build":
		graft.ExecuteHandlers("build")
	case "test":
		graft.ExecuteHandlers("test")
	case "help":
		logger.Warning("TODO")
	default:
		logger.Error("Command %s does not exist, run `graft help` to see the list of commands", action)
	}
}
