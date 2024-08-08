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

const (
	// DirPerm defines the permission for created directories
	DirPerm = 0755
)

func CompileAndExecuteGraft(graftFile string, action string) error {
	// Create bin directory if it doesn't exist
	binDir := filepath.Join(filepath.Dir(graftFile), "bin")
	if err := os.MkdirAll(binDir, DirPerm); err != nil {
		return fmt.Errorf("failed to create bin directory: %w", err)
	}
	
	executableName := "graft_config"
	if runtime.GOOS == "windows" {
		executableName += ".exe"
	}
	executablePath := filepath.Join(binDir, executableName)
	
	cmd := exec.Command("go", "build", "-tags", "graft", "-o", executablePath, graftFile)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to compile graft: %w", err)
	}
	
	logger.Success(`[Graft]: Successfully compiled "graft.go", running it...`)
	
	cmd = exec.Command(executablePath, action)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Dir = filepath.Dir(graftFile)
	
	return cmd.Run()
}
