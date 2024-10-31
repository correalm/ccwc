package main

import (
	"ccwc/cmd/cli/command"
	"ccwc/internal/interfaces"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	args_len := len(args)

	var filename string

	if args_len > 0 {
		filename = args[args_len-1]
	}

	cmds := []interfaces.Command{
		command.NewCountLines(filename),
		command.NewCountWords(filename),
		command.NewByteCounter(filename),
	}

	parser := command.NewParser(cmds)

	if err := parser.Parse(args, command.Help); err != nil {
		os.Stderr.WriteString(fmt.Sprintf("error: %v", err.Error()))
		os.Exit(1)
	}
}
