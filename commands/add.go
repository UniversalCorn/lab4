package engine

import (
	"fmt"
	"strconv"
)

type Add struct {
	Arg []string
}

func (a *Add) Execute(h Handler) {
	if len(a.Arg) < 1 {
		h.Post(&PrintCommand{Arg: "error: expected amount of arguments doesn't match provided [add]"}, true)
		return
	}

	res := 0
	for _, i := range a.Arg {
		if number, err := strconv.Atoi(i); err != nil {
			h.Post(&PrintCommand{Arg: fmt.Sprintf("error: argument is not a number [add] - %s", err)}, true)
		} else {
			res += number
		}
	}

	h.Post(&PrintCommand{Arg: strconv.Itoa(res)}, true)
}
