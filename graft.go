//go:build graft
// +build graft

package main

import (
	"github.com/TomasBorquez/graft/pkg"
)

func main() {
	graft.ExecuteTasks(func(t *graft.TaskExecutor) {
		t.DefineTask("build", func(p *graft.TaskConfig) {
			p.Build(graft.BuildOptions{
				OutputPath: "bin/graft",
				SourcePath: "cmd/graft/main.go",
			})
		})
		
		startConfig := graft.HRConfig{
			Action:       "start",
			ExcludeDir:   []string{".git", ".idea", "documentation", "images", "bin"},
			IncludeExten: []string{".go", ".mod"},
		}
		t.DefineHotReloadTask(startConfig, func(p *graft.TaskConfig) {
			p.Run("graft", "build")
		})
	})
}
