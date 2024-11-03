package command

import "fmt"

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
