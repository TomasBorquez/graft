package main

import (
	"github.com/TomasBorquez/graft/internal/runner"
	"github.com/TomasBorquez/graft/internal/scripts"
	"github.com/TomasBorquez/logger"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		scripts.PrintHelp()
		return
	}

	action := args[0]

	switch action {
	case "help":
		scripts.PrintHelp()
		return
	case "init":
		scripts.InitGraft()
		return
	case "start":
	case "build":
	case "test":
	case "format":
	case "lint":
	default:
		logger.Error("[Graft]: Command %s does not exist, run `graft help` to see the list of commands", action)
		return
	}

	file, err := runner.FindGraftFile()
	if err != nil {
		logger.Error(`[Graft]: Failed to find "graft.go" in root directory: %v`, err)
		return
	}

	logger.Success(`[Graft]: Found "graft.go", compiling and executing...`)
	err = runner.CompileAndExecuteGraft(file, action)

	if err != nil {
		logger.Error(`[Graft]: Error meanwhile compiling "graft.go" %v`, err)
		return
	}
}
