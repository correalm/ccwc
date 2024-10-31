package command

import (
	"ccwc/internal/interfaces"
)

type Parser struct {
	commands []interfaces.Command
}

type helper func()

func NewParser(commands []interfaces.Command) *Parser {
	return &Parser{commands: commands}
}

func (p *Parser) Parse(args []string, helper helper) error {
	args_len := len(args)

	if args_len < 1 {
    helper()
		return nil
	}

	for _, cmd := range p.commands {
		if args_len > 1 {
			for i := 0; i < args_len-1; i++ {
				arg := args[i]

				if cmd.Name() == arg {
					cmd.ParseFlags(args[i:i])
					cmd.Run()
				}
			}
		} else {
			cmd.ParseFlags(args[1:])
			cmd.Run()
		}
	}

	return nil
}
