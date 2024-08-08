package graft

import (
	"fmt"
	"github.com/TomasBorquez/logger"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Project struct {
	name string
}

type Callback = func(p *Project)

func Config(name string, cb Callback) {
	if len(os.Args) <= 1 {
		logger.Error("[Graft]: Use `graft action` to execute this file")
		return
	}

	action := os.Args[1]
	if name == action {
		cb(&Project{name: action})
		if name == "build" {
			logger.Success("[Graft]: Successfully built project")
		}
	}
}

// Run runs what you give it in the commandline
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

// BuildOptions contains the configuration for building a project.
type BuildOptions struct {
	OutputPath string   // The path where the compiled binary will be saved
	SourcePath string   // The path to the main package to compile (optional)
	Tags       []string // Build tags to be passed to the compiler
	LDFlags    []string // Linker flags to be passed to the compiler
}

// Build compiles the project with the specified options.
func (p *Project) Build(opts BuildOptions) {
	if opts.OutputPath == "" {
		opts.OutputPath = p.name
	}

	if runtime.GOOS == "windows" && !strings.HasSuffix(opts.OutputPath, ".exe") {
		opts.OutputPath += ".exe"
	}

	args := []string{"build", "-o", opts.OutputPath}

	if len(opts.Tags) > 0 {
		args = append(args, "-tags", strings.Join(opts.Tags, ","))
	}

	if len(opts.LDFlags) > 0 {
		args = append(args, "-ldflags", strings.Join(opts.LDFlags, " "))
	}

	if opts.SourcePath != "" {
		args = append(args, opts.SourcePath)
	}

	p.Run("go", args...)
}
