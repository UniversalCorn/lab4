package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"

	engine "github.com/UniversalCorn/lab4/commands"
	"github.com/UniversalCorn/lab4/eventloop"
	"github.com/UniversalCorn/lab4/parser"
)

func main() {

	file, err := os.OpenFile("input", os.O_RDONLY, os.ModePerm)
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return
	}

	defer file.Close()

	var eventLoop eventloop.EventLoop
	eventLoop.Start()

	reader := bufio.NewReader(file)

	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("read line error: %v", err)
			return
		}
		command := parser.Parse((string(line)))
		eventLoop.Post(command, false)

	}

	eventLoop.AwaitFinish()
	eventLoop.Post(&engine.PrintCommand{Arg: "Should end!"}, false)
	time.Sleep(10 * time.Second)
}
