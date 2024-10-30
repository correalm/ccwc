package command

import "fmt"

func help() {
	help := "usage: [<flags>] filename"
	fmt.Println(help)
}
