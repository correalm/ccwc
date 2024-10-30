package command

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

type ByteCounter struct {
  fs *flag.FlagSet
  filename string
}

func NewByteCounter(filename string) *ByteCounter {
  cmd := &ByteCounter{
    fs: flag.NewFlagSet("-b", flag.ContinueOnError),
    filename: filename,
  }

  return cmd
}

func (cmd *ByteCounter) Name() string {
  return cmd.fs.Name()
}

func (cmd *ByteCounter) ParseFlags(flags []string) error {
  if len(flags) == 0 { return fmt.Errorf("missing flags") }

  return cmd.fs.Parse(flags)
}

func (cmd *ByteCounter) Run() error {
	file, err := os.Open(cmd.filename)
	defer file.Close()

	if err != nil { return err }

  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanBytes)

  counter := 0

  for scanner.Scan() { counter++ }
  fmt.Println("Bytes:", counter)
  return err
}
