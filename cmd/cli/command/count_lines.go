package command

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type LineCounter struct {
  fs *flag.FlagSet
  filename string
}

func NewCountLines(filename string) *LineCounter {
  cmd := &LineCounter{
    fs: flag.NewFlagSet("-l", flag.ContinueOnError),
    filename: filename,
  }

  return cmd
}

func (cmd *LineCounter) Name() string {
  return cmd.fs.Name()
}

func (cmd *LineCounter) ParseFlags(flags []string) error {
  if len(flags) == 0 { return fmt.Errorf("missing flags") }

  return cmd.fs.Parse(flags)
}

func (cmd *LineCounter) Run() error {
	file, err := os.Open(cmd.filename)
	defer file.Close()

	if err != nil { return err }

  scanner := bufio.NewScanner(file)
  counter := 0

  for scanner.Scan() { counter++ }
  fmt.Println("Lines:", counter)
  return err
}
