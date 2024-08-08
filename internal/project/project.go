package graft

import (
	"errors"
	"fmt"
	"github.com/TomasBorquez/graft/internal/store"
	"github.com/TomasBorquez/logger"
	"os/exec"
)

type Project struct {
	name string
}

func (p *Project) Run(name string, args ...string) {
	graft.PushHandlers(name, func() error {
		if len(args) == 0 {
			logger.Error("Error: No command provided")
			return errors.New("no command provided")
		}
		
		cmd := exec.Command(name, args...)
		output, err := cmd.CombinedOutput()
		
		if err != nil {
			return err
		}
		
		fmt.Println(string(output))
		return nil
	})
}
