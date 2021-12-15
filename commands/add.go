package commands

import (
	"fmt"
	"strconv"

	"github.com/UniversalCorn/lab4/eventloop"
)

type Add struct {
	Args []string
}

func (a *Add) Execute(h eventloop.Handler) {
	if len(a.Args) < 1 {
		h.Post(&Println{Message: "error: expected amount of arguments doesn't match provided [add]"}, true)
		return
	}

	res := 0
	for _, i := range a.Args {
		if number, err := strconv.Atoi(i); err != nil {
			h.Post(&Println{Message: fmt.Sprintf("error: argument is not a number [add] - %s", err)}, true)
		} else {
			res += number
		}
	}

	h.Post(&Println{Message: strconv.Itoa(res)}, true)
}
