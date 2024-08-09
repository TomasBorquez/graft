package scripts

import (
	"github.com/TomasBorquez/logger"
	"os"
)

func InitGraft() {
	content := `//go:build graft
// +build graft

package main

import (
	"github.com/TomasBorquez/graft/pkg"
)

func main() {
  graft.ExecuteTasks(func(t *graft.TaskExecutor) {
    t.DefineTask("build", func(p *graft.TaskConfig) {
      p.Build(graft.BuildOptions{
        OutputPath: "bin/app",
        SourcePath: "./main.go",
      })
    })

		var startConfig = graft.HRConfig{
			Action: "start",
		}
    t.DefineHotReloadTask(startConfig, func(p *graft.TaskConfig) {
      p.Run("go", "run", "main.go")
    })
  })
}`

	if _, err := os.Stat("graft.go"); err == nil {
		logger.Warning("[Graft]: `graft.go` already exists. Skipping creation.")
		return
	}

	file, err := os.Create("graft.go")
	if err != nil {
		logger.Error("[Graft]: Error creating `graft.go`: %v", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		logger.Error("[Graft]: Error writing to `graft.go: %v", err)
		return
	}

	logger.Success("[Graft]: Successfully created `graft.go`")
	return
}

func PrintHelp() {
	logger.Custom("%sGraft%s is a build tool for Go that allows you to make your own build scripts in Go.", logger.Blue, logger.Reset)
	logger.Custom("")
	logger.Custom("Usage: graft %s<command>%s [...args]", logger.Blue, logger.Reset)
	logger.Custom("")
	logger.Custom("Commands:")
	logger.Custom("   " + logger.Orange + "init" + logger.Reset + "          It adds a simple template `graft.go` file.")
	logger.Custom("   " + logger.Orange + "format" + logger.Reset + "        Executes the format script but if it doesn't exist it formats all the files.")
	logger.Custom("   " + logger.Orange + "help" + logger.Reset + "          Shows this menu.")
	logger.Custom("")
	logger.Custom("   " + logger.Red + "start" + logger.Reset + "         Executes the start script.")
	logger.Custom("   " + logger.Red + "dev" + logger.Reset + "           Executes the dev script.")
	logger.Custom("   " + logger.Red + "build" + logger.Reset + "         Executes the build script.")
	logger.Custom("   " + logger.Red + "test" + logger.Reset + "          Executes the test script.")
	logger.Custom("   " + logger.Red + "lint" + logger.Reset + "          Executes the lint script.")
}
