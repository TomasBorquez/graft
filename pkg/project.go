package graft

import (
	"fmt"
	"github.com/TomasBorquez/logger"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

type Project struct {
	name string
	Path string
}

type Callback = func(p *Project)

func Config(name string, cb Callback) {
	if len(os.Args) <= 1 {
		logger.Error("[Graft]: Use `graft action` to execute this file")
		return
	}

	action := os.Args[1]
	path, err := os.Getwd()
	if err != nil {
		logger.Error("[Graft]: Error meanwhile getting working directory: %v", err)
		return
	}

	if name == action {
		cb(&Project{name: action, Path: path})
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

	fmt.Printf("[Graft]: Running command `%s %v`\n", command, args)
	cmd := exec.Command(command, args...)
	cmd.Dir = p.Path

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("%v", err)
		os.Exit(1)
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

// FormatOptions contains the configuration for formatting options
type FormatOptions struct {
	File       string // Specific file to format, empty means format all
	Simplified bool   // Whether to use simplified formatting (-s flag)
	Write      bool   // Whether to write changes to file(s) (-w flag)
	List       bool   // Whether to list files that would be formatted (-l flag)
}

// Format applies Go formatting to the project files based on the given options
func (p *Project) Format(opts FormatOptions) {
	var args []string

	cmd := "gofmt"
	if opts.Simplified {
		args = append(args, "-s")
	}

	if opts.Write {
		args = append(args, "-w")
	}

	if opts.List {
		args = append(args, "-l")
	}

	if opts.File != "" {
		args = append(args, opts.File)
	} else {
		args = append(args, "./..")
	}

	p.Run(cmd, args...)
	return
}

// FormatAll formats all Go files in the project with simplified formatting
func (p Project) FormatAll() {
	p.Format(FormatOptions{
		Simplified: true,
		Write:      true,
	})
	return
}

// ListUnformatted lists all Go files in the project that need formatting
func (p Project) ListUnformatted() {
	p.Format(FormatOptions{
		Simplified: true,
		List:       true,
	})
	return
}
