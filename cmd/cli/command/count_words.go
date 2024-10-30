package command

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type WordCounter struct {
  fs *flag.FlagSet
  filename string
}

func NewCountWords(filename string) *WordCounter {
  cmd := &WordCounter{
    fs: flag.NewFlagSet("-w", flag.ContinueOnError),
    filename: filename,
  }

  return cmd
}

func (cmd *WordCounter) Name() string {
  return cmd.fs.Name()
}

func (cmd *WordCounter) ParseFlags(flags []string) error {
  if len(flags) == 0 { return fmt.Errorf("missing flags") }

  return cmd.fs.Parse(flags)
}

func (cmd *WordCounter) Run() error {
	file, err := os.Open(cmd.filename)
	defer file.Close()

	if err != nil { return err }

  scanner := bufio.NewScanner(file)

  scanner.Split(bufio.ScanWords)
  counter := 0

  for scanner.Scan() { counter++ }
  fmt.Println("Words:", counter)
  return err
}
