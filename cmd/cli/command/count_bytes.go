package command

import (
	"bufio"
	"ccwc/services"
	"flag"
	"fmt"
)

type ByteCounter struct {
	fs       *flag.FlagSet
	filename string
}

func NewByteCounter(filename string) *ByteCounter {
	cmd := &ByteCounter{
		fs:       flag.NewFlagSet("-b", flag.ContinueOnError),
		filename: filename,
	}

	return cmd
}

func (cmd *ByteCounter) Name() string {
	return cmd.fs.Name()
}

func (cmd *ByteCounter) ParseFlags(flags []string) error {
	if len(flags) == 0 {
		return fmt.Errorf("missing flags")
	}

	return cmd.fs.Parse(flags)
}

func (cmd *ByteCounter) Run() error {
	message, err := services.Counter(cmd.filename, "Bytes:", bufio.ScanBytes)

	if err != nil {
		return err
	}

	fmt.Println(message)

	return nil
}
