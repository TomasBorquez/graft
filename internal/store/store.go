package graft

import (
	"github.com/TomasBorquez/logger"
	"os"
)

type handlerCb = func() error

var handlers = make(map[string][]handlerCb, 4)

func PushHandlers(name string, cb handlerCb) {
	handlers[name] = append(handlers[name], cb)
}

func ExecuteHandlers(name string) {
	currHandler := handlers[name]
	if currHandler == nil {
		logger.Error("[Graft]: You still haven't set up your %s script", name)
		logger.Warning("[Graft]: Stopping execution...")
		os.Exit(0)
	}
	
	for _, handler := range currHandler {
		err := handler()
		logger.Error("[Graft]: Error meanwhile executing `graft %s`", name)
		logger.Error("%v", err)
		logger.Warning("[Graft]: Stopping execution...")
		os.Exit(0)
	}
}