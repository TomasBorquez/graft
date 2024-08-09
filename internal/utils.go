package internal

import (
	"fmt"
	"github.com/TomasBorquez/logger"
	"os"
	"os/exec"
)

func Cmd(command string, args ...string) {
	if len(args) == 0 {
		logger.Error("[Graft]: No args provided on `Run(command, args)`")
		return
	}

	PrintCommand(command, args...)

	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		logger.Error("%v", err)
	}
}

func PrintCommand(command string, args ...string) {
	fmt.Printf("[Graft]: Running command `%s", command)
	for _, arg := range args {
		fmt.Printf(" %s", arg)
	}
	fmt.Println("`")
}
