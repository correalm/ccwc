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

  if args_len == 0 {
    os.Stderr.WriteString(fmt.Sprint("No args provided\n"))
    command.Help()
    os.Exit(1)
  }

  path := args[args_len-1]

  stat, err := os.Stat(path)

  if os.IsNotExist(err) {
    os.Stderr.WriteString(fmt.Sprintf("error: %v\n", err))
    os.Exit(1)
  }

  if stat.IsDir() {
    os.Stderr.WriteString(fmt.Sprintf("error: %s is a directory\n", stat.Name()))
    os.Exit(1)
  }

  cmds := []interfaces.Command{
    command.NewLineCounter(path),
    command.NewWordCounter(path),
    command.NewByteCounter(path),
    command.NewHelp(),
  }

  parser := command.NewParser(cmds)

  if err := parser.Parse(args, command.Help); err != nil {
    os.Stderr.WriteString(fmt.Sprintf("error: %v", err.Error()))
    os.Exit(1)
  }
}
