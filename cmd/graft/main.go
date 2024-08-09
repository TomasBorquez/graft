package main

import (
	"github.com/TomasBorquez/graft/internal/runner"
	"github.com/TomasBorquez/graft/internal/scripts"
	"github.com/TomasBorquez/logger"
	"os"
)

func main() {
	if len(os.Args) == 0 {
		scripts.PrintHelp()
		return
	}

	action := os.Args[1]

	switch action {
	case "help":
		scripts.PrintHelp()
		return
	case "init":
		scripts.InitGraft()
		return
	case "start", "dev", "build", "test", "format", "lint":
		executeConfig(action)
		return
	default:
		logger.Error("[Graft]: Command %s does not exist, run `graft help` to see the list of commands", action)
		return
	}
}

func executeConfig(action string) {
	file, err := runner.FindGraftFile()
	if err != nil {
		logger.Error(`[Graft]: Failed to find "graft.go" in root directory: %v`, err)
		return
	}

	logger.Custom(`[Graft]: Found "graft.go", compiling and executing...`)
	err = runner.CompileAndExecuteGraft(file, action)

	if err != nil {
		logger.Error(`[Graft]: Error meanwhile compiling "graft.go" %v`, err)
		return
	}
}
