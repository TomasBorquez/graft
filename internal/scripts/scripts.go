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
	graft.Config("build", func(p *graft.Project) {
		p.Build(graft.BuildOptions{
			OutputPath: "bin/app.exe",
			SourcePath: "./main.go",
		})
	})

	graft.Config("start", func(p *graft.Project) {
		p.Run("go", "run", "main.go")
	})

	graft.Config("lint", func(p *graft.Project) {
		p.FormatAll()
	})
}`
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
	logger.Warning("TODO")
}
