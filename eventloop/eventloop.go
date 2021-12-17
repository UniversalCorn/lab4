package eventloop

import (
	"fmt"

	engine "github.com/UniversalCorn/lab4/commands"
)

type EventLoop struct {
	queue      *Queue
	isStop     bool
	stopSignal chan bool
}

func (e *EventLoop) Start() {
	/*
		just for debug
	*/
	fmt.Println("EVENTLOOP STARTED")
	e.queue = &Queue{}
	e.stopSignal = make(chan bool, 1)

	go func() {
		for !e.isStop || !e.queue.empty() {
			command := e.queue.pop()
			command.Execute(e)
		}
		e.stopSignal <- true
	}()
}

func (e *EventLoop) Post(command engine.Command, isInner bool) {
	if !e.isStop {
		e.queue.push(command)
	} else if isInner {
		e.Start()
		e.queue.push(command)
		e.AwaitFinish()
	}
}

func (e *EventLoop) AwaitFinish() {
	e.Post(engine.CommandFunc(func(h engine.Handler) {
		e.isStop = true
	}), false)
	<-e.stopSignal
	/*
		just for debug
	*/
	fmt.Println("EVENTLOOP STOPPED")
}
