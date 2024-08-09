//go:build graft
// +build graft

package main

import (
	"github.com/TomasBorquez/graft/pkg"
)

func main() {
	graft.Config("build", func(p *graft.Project) {
		p.Build(graft.BuildOptions{
			OutputPath: "bin/graft.exe",
			SourcePath: "cmd/graft/main.go",
		})
	})

	graft.Config("format", func(p *graft.Project) {
		p.FormatAll()
	})
}
