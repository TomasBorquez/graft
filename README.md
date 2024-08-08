<div align="center">
	<img alt="Graft logo" src="/images/logo.svg" height="300" /><br />
</div>

<hr>

<div align="center">
  <b>Graft</b> is a powerful Go library designed to streamline and enhance the development workflow for Go projects. It
  provides a set of tools and utilities to simplify common tasks, automate build processes, and improve overall
  productivity.
</div>

### Features

- **Customizable Build Configuration**: Easily configure and manage your build process with custom options and tags.
- **Hot Reloading**: Enable hot reloading for faster development cycles.
- **Integrated Testing**: Run tests with coverage reporting and verbose output.
- **Linting**: Perform code linting with optional auto-fixing.
- **Benchmarking**: Run benchmarks with customizable options.
- **Environment Variable Management**: Easily set environment variables from .env files.
- **Version Management**: Automate version bumping, including updating files, creating git tags, and generating
  changelogs.
- **Custom Task Support**: Define and run custom tasks tailored to your project's needs.

### Installation

To install the command line tool use:
```shell
go install github.com/TomasBorquez/graft
```

To add graft to your project just use:
```shell
go get github.com/TomasBorquez/graft
```

### Quick Start

1. Use `graft init` to or create the `graft.go` file manually in the **root directory**:
  ```go
  //go:build graft
  // +build graft

  package main

  import (
    "github.com/TomasBorquez/graft/pkg"
  )

  func main() {
    graft.Config("build", func(p *graft.Project) {
      p.Build(graft.BuildOptions{
        OutputPath: "bin/myapp",
        SourcePath: "cmd/myapp/main.go",
      })
    })
  }
  ```
2. Run `graft build` and that's it!

No `.sh` nor `.bat` files no `make` files, just some simple go code that makes your build compatible with multiple operating systems out of the box.

### Documentation
[TODO](TODO.md)

### Contributing
[TODO](TODO.md)

### License
Graft is released under the MIT License. See the [LICENSE](LICENSE) file for more details.