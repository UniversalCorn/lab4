package commands

import (
	"fmt"

	"github.com/UniversalCorn/lab4/eventloop"
)

type Println struct {
	Message string
}

func (p *Println) Execute(_ eventloop.Handler) {
	fmt.Println(p.Message)
}
