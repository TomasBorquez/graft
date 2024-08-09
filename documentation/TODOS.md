## Planned stuff

- [ ] Features
    - [x] Basic commands:
        - [x] `graft init`
        - [x] `graft help`
    - [ ] Abstractions
        - [x] Run formater `p.Format()`
        - [ ] Test command `p.Test()`
        - [ ] .env variable setter `p.Env()`
        - [ ] go mod download/tidy `p.Module()`
        - [ ] go update dependencies `p.ManageDependencies()`
        - [ ] Docker command `p.DockerCompose()`
        - [ ] Run linters `p.RunLinter()`
    - [ ] Hot reloading option
        - [ ] Default config
        - [ ] `WatchDir`
        - [ ] `Ignore`
        - [ ] `Command`
    - [ ] Log how much time it took to run each script
    - [ ] Benchmark script/functions
    - [ ] Deploy functions
    - [ ] Functioning `go get graft`
    - [ ] Print graft version on `graft help`
- [ ] Extras
    - [ ] Fix go report
    - [ ] Add git hooks
        - [ ] gofmt on commit
        - [ ] golint on push
    - [ ] Testing
        - [ ] Unit tests for each `/pkg` function
        - [ ] Integration tests for `/internal` functions
        - [ ] Add CI/CD
        - [ ] Add goreportcard on readme about passing build
    - [ ] Documentation
        - [x] Basic `README.md` examples
        - [ ] `pkg.go` example functions for all of `/pkg`
        - [ ] Website examples
        - [ ] Make a few videos showcasing
        - [ ] Spread the word 🐣