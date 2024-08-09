package graft

import (
	"github.com/TomasBorquez/graft/internal"
	"github.com/TomasBorquez/logger"
	"os"
)

type TaskExecutor struct{}

func ExecuteTasks(cb func(t *TaskExecutor)) {
	if len(os.Args) <= 1 {
		logger.Error("[Graft]: Use `graft action` to execute this file")
		return
	}
	
	store.action = os.Args[1]
	
	path, err := os.Getwd()
	if err != nil {
		logger.Error("[Graft]: Error meanwhile getting working directory: %v", err)
		return
	}
	
	store.path = path
	
	cb(&TaskExecutor{})
	
	if !store.executedAction {
		if store.action == "format" {
			logger.Warning("[Graft]: No format task found, executing default formatter...")
			internal.Cmd("gofmt", "-s", "-w", "./..")
			return
		}
		
		logger.Error("[Graft]: No script found for `graft %s`", store.action)
		return
	}
}

type TaskConfig struct{}
type TaskCb = func(p *TaskConfig)

func (t *TaskExecutor) DefineTask(name string, cb TaskCb) {
	if name == store.action {
		store.executedAction = true
		cb(&TaskConfig{})
		
		if name == "build" {
			logger.Success("[Graft]: Successfully built project")
		}
	}
}

type HRConfig struct {
	Action       string
	IncludeDir   []string
	ExcludeDir   []string
	IncludeExten []string
	ExcludeExten []string
}

func (t *TaskExecutor) DefineHotReloadTask(opts HRConfig, cb TaskCb) {
	if opts.Action == store.action {
		store.executedAction = true
		
		AddWatcher(&opts, cb)
		
		if opts.Action == "build" {
			logger.Success("[Graft]: Successfully built project")
		}
	}
}
