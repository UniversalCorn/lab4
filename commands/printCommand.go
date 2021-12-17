package engine

import (
	"fmt"
)

type PrintCommand struct {
	Arg string
}

func (p *PrintCommand) Execute(_ Handler) {
	fmt.Println(p.Arg)
}
