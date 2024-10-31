package command

import "fmt"

func Help() {
	help := "usage: [<flags>] filename"
	fmt.Println(help)
}
