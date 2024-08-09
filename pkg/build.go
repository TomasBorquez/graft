package graft

import (
	"github.com/TomasBorquez/graft/internal"
	"runtime"
	"strings"
)

// BuildOptions contains the configuration for building a project.
type BuildOptions struct {
	OutputPath string   // The path where the compiled binary will be saved
	SourcePath string   // The path to the main package to compile (optional)
	Tags       []string // Build tags to be passed to the compiler
	LDFlags    []string // Linker flags to be passed to the compiler
}

// Build compiles the project with the specified options.
func (p *TaskConfig) Build(opts BuildOptions) {
	if opts.OutputPath == "" {
		opts.OutputPath = store.action
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

// Run runs what you give it in the commandline
func (p *TaskConfig) Run(command string, args ...string) {
	internal.Cmd(command, args...)
}
