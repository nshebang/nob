package main

import (
	"os"
)

type Runner struct {
	commands	map[string]*Command	
}

func NewRunner() *Runner {
	return &Runner {
		commands: make(map[string]*Command),
	}
}

func (r *Runner) Execute() int {
	if len(os.Args) < 2 {
		PrintHelp()
		return 1
	}

	cmdname := os.Args[1]
	params := os.Args[2:]
 
	cmd := r.commands[cmdname]

	if cmd == nil {
		PrintHelp()
		return 1
	}

	return cmd.Run(&params)
}

func (r *Runner) AppendCmd(name string, command *Command) {
	r.commands[name] = command
}

