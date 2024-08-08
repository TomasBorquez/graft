package graft

import (
	"fmt"
	"github.com/TomasBorquez/logger"
	"os"
	"os/exec"
)

type Project struct {
	name string
}

func Config(name string, cb func(g *Project)) {
	action := os.Args[1]
	if name == action {
		cb(&Project{name: action})
	}
}

func (p *Project) Run(command string, args ...string) {
	if len(args) == 0 {
		logger.Error("[Graft]: Error meanwhile running `graft %s`", p.name)
		logger.Error("No command provided on `g.Run(command, args)`")
		return
	}
	
	cmd := exec.Command(command, args...)
	output, err := cmd.CombinedOutput()
	
	if err != nil {
		return
	}
	
	fmt.Println(string(output))
}
