package parser

import (
	"strings"

	engine "github.com/UniversalCorn/lab4/commands"
)

func Parse(line string) engine.Command {

	values := strings.Fields(strings.TrimSpace(line))
	if len(values) < 1 {
		return &engine.PrintCommand{Arg: "error: blank line"}
	}

	command := values[0]
	args := values[1:]

	if command == "add" {
		return &engine.Add{Arg: args}
	} else if command == "print" {
		return &engine.PrintCommand{Arg: strings.Join(args, " ")}
	} else {
		return &engine.PrintCommand{Arg: "error: unexpected method name"}
	}
}
