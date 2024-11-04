package command

import (
	"bufio"
	"ccwc/services"
	"flag"
	"fmt"
)

type LineCounter struct {
	fs       *flag.FlagSet
	filename string
}

func NewLineCounter(filename string) *LineCounter {
	cmd := &LineCounter{
		fs:       flag.NewFlagSet("-l", flag.ContinueOnError),
		filename: filename,
	}

	return cmd
}

func (cmd *LineCounter) Name() string {
	return cmd.fs.Name()
}

func (cmd *LineCounter) ParseFlags(flags []string) error {
	if len(flags) == 0 {
		return fmt.Errorf("missing flags")
	}

	return cmd.fs.Parse(flags)
}

func (cmd *LineCounter) Run() error {
	message, err := services.Counter(cmd.filename, "Lines:", bufio.ScanLines)

	if err != nil {
		return err
	}

	fmt.Println(message)

	return nil
}
