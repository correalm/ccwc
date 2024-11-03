package command

import (
	"flag"
	"fmt"
)

type Helper struct {
  fs *flag.FlagSet
}

func NewHelp() *Helper {
  cmd := &Helper{
    fs: flag.NewFlagSet("-h", flag.ContinueOnError),
  }

  return cmd
}

func (cmd *Helper) Name() string {
  return cmd.fs.Name()
}

func (cmd *Helper) ParseFlags(flags []string) error {
  if len(flags) == 0 {
    return fmt.Errorf("missing flags")
  }

  return cmd.fs.Parse(flags)
}

func (cmd *Helper) Run() error {
  Help()

  return nil
}

func Help() {
	help := `usage: ./ccwc [<flags>] filename

  Available commands:
  -l print the number of lines
  -b print the number of bytes
  -w print the number of words

  More than one command:
  ./ccwc -l -w filename

  When no flags are provided, the program will print all information
`
	fmt.Println(help)
}
