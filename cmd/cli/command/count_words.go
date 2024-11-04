package command

import (
	"bufio"
	"ccwc/services"
	"flag"
	"fmt"
)

type WordCounter struct {
	fs       *flag.FlagSet
	filename string
}

func NewWordCounter(filename string) *WordCounter {
	cmd := &WordCounter{
		fs:       flag.NewFlagSet("-w", flag.ContinueOnError),
		filename: filename,
	}

	return cmd
}

func (cmd *WordCounter) Name() string {
	return cmd.fs.Name()
}

func (cmd *WordCounter) ParseFlags(flags []string) error {
	if len(flags) == 0 {
		return fmt.Errorf("missing flags")
	}

	return cmd.fs.Parse(flags)
}

func (cmd *WordCounter) Run() error {
	message, err := services.Counter(cmd.filename, "Words:", bufio.ScanWords)

	if err != nil {
		return err
	}

	fmt.Println(message)

	return nil
}
