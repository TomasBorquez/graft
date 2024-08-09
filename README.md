<div align="center">
	<img alt="Graft logo" src="/images/logo.svg" height="400" /><br />
    <a href="https://pkg.go.dev/github.com/TomasBorquez/graft">
      <img src="https://img.shields.io/badge/%F0%9F%93%9A%20godoc-pkg-00ACD7.svg?color=00ACD7&style=flat-square">
    </a>
    <a href="https://goreportcard.com/report/github.com/TomasBorquez/graft">
      <img src="https://goreportcard.com/badge/github.com/TomasBorquez/graft">
    </a>
</div>

<hr>

<div align="center">
  <b>Graft</b> is a Go build library in Go for building Go projects :) it
  provides a set of tools and utilities to simplify common build tasks, 
  automate build processes, hot reloading and more.
</div>

<hr>

### 🦟 Features

- **Customizable Build Configuration**: Easily configure and manage your build process with custom options and tags.
- **Hot Reloading**: Enable hot reloading for faster development cycles.
- **Integrated Testing**: Run tests with coverage reporting and verbose output.
- **Linting**: Perform code linting with optional auto-fixing.
- **Benchmarking**: Run benchmarks with customizable options.
- **Environment Variable Management**: Easily set environment variables from .env files.
- **Version Management**: Automate version bumping, including updating files, creating git tags, and generating
  changelogs.
- **Custom Task Support**: Define and run custom tasks tailored to your project's needs.

<hr>

### ⚙ Installation

To install the command line tool use:

```shell
go install github.com/TomasBorquez/graft
```

To add graft to your project just use:

```shell
go get github.com/TomasBorquez/graft
```

<hr>

### ⚡ Quick Start

1. Use `graft init` to or create the `graft.go` file manually in the **root directory**:

```go
//go:build graft
// +build graft

package main

import (
  "github.com/TomasBorquez/graft/pkg"
)

func main() {
  graft.ExecuteTasks(func(t *graft.TaskExecutor) {
    t.DefineTask("build", func(p *graft.TaskConfig) {
      p.Build(graft.BuildOptions{
        OutputPath: "bin/graft",
        SourcePath: "cmd/graft/main.go",
      })
    })

    t.DefineHotReloadTask(graft.HotReloadingConfig{}, func(p *graft.TaskConfig) {
      p.Run("go", "run", "main.go")
    })

    t.DefineTask("format", func(p *graft.TaskConfig) {
      p.FormatAll()
    })
  })
}
```

2. Run `graft build` and that's it!

No `.sh` nor `.bat` files no `make` files, just some simple go code that makes your build compatible with multiple
operating systems out of the box.

<hr>

### 🤔 Why?

As a [zig](https://ziglang.org/) lover, when I got into Go I wondered why isn't there a native tool for building your Go
project in **Go itself** such as [zig build](https://ziglang.org/learn/build-system/), you look at any complex go project,
and it's full of `.sh` or `make` files (and `.bat` files if they thought about the Windows users). It didn't really feel right, 
this is why I created **graft**, so I can start a project, add some simple script **in go** and focus on programming and not in setting my build process.

<hr>

### ⛪ Contributing

[TODO](TODO.md)

<hr>

### License

Graft is released under the MIT License. See the [LICENSE](LICENSE) file for more details.
