package runner

import (
	"fmt"
	"github.com/TomasBorquez/logger"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func FindGraftFile() (string, error) {
	currentDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	
	for {
		graftFile := filepath.Join(currentDir, "graft.go")
		if _, err := os.Stat(graftFile); err == nil {
			return graftFile, nil
		}
		
		parentDir := filepath.Dir(currentDir)
		if parentDir == currentDir {
			return "", fmt.Errorf("graft.go not found in project hierarchy")
		}
		currentDir = parentDir
	}
}

func CompileAndExecuteGraft(graftFile string, action string) error {
	tempDir, err := os.MkdirTemp("", "graft")
	if err != nil {
		return err
	}
	defer os.RemoveAll(tempDir)
	
	executableName := "graft_config"
	if runtime.GOOS == "windows" {
		executableName += ".exe"
	}
	executablePath := filepath.Join(tempDir, executableName)
	
	cmd := exec.Command("go", "build", "-o", executablePath, graftFile)
	if err := cmd.Run(); err != nil {
		return err
	}
	
	logger.Success(`[Graft]: Successfully compiled "graft.go", running it`)
	
	cmd = exec.Command(executablePath, action)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = filepath.Dir(graftFile)
	
	return cmd.Run()
}
