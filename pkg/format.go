package graft

// FormatOptions contains the configuration for formatting options
type FormatOptions struct {
	File       string // Specific file to format, empty means format all
	Simplified bool   // Whether to use simplified formatting (-s flag)
	Write      bool   // Whether to write changes to file(s) (-w flag)
	List       bool   // Whether to list files that would be formatted (-l flag)
}

// Format applies Go formatting to the project files based on the given options
func (p *TaskConfig) Format(opts FormatOptions) {
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
func (p *TaskConfig) FormatAll() {
	p.Format(FormatOptions{
		Simplified: true,
		Write:      true,
	})
	return
}

// ListUnformatted lists all Go files in the project that need formatting
func (p *TaskConfig) ListUnformatted() {
	p.Format(FormatOptions{
		Simplified: true,
		List:       true,
	})
	return
}
