package graft

type Store struct {
	path           string
	executedAction bool
	action         string
}

var store = Store{
	path:           "",
	action:         "",
	executedAction: false,
}
