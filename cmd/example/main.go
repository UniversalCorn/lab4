package main

import (
	engine "github.com/UniversalCorn/lab4/commands"
	"github.com/UniversalCorn/lab4/eventloop"
)

func main() {
	// read file with commnads
	// take all lines
	//and pass to Parse function
	var eventLoop eventloop.EventLoop
	eventLoop.Start()
	eventLoop.AwaitFinish()
	eventLoop.Post(&engine.PrintCommand{Arg: "Should end!"}, false)
}
