package eventloop

import "fmt"

type EventLoop struct {
	queue *Queue

	stop bool
	stopSignal chan struct{}
}

func (e *EventLoop) Start() {
	/*
		just for debug
	*/
	fmt.Println("EVENTLOOP STARTED")
	e.queue = &Queue{ signal: make(chan struct{}) }
	e.stopSignal = make(chan struct{})

	go func() {
		for !e.stop || !e.queue.empty() {
			command := e.queue.pop()
			command.Execute(e)
		}
		e.stopSignal <- struct{}{}
	}()
}

func (e *EventLoop) Post(command Command, isInner bool) {
	if !e.stop {
		e.queue.push(command)
	} else if isInner {
		e.Start()
		e.queue.push(command)
		e.AwaitFinish()
	}
}

func (e *EventLoop) AwaitFinish() {
	e.Post(CommandFunc(func(h Handler) {
		e.stop = true
	}), false)
	<- e.stopSignal
	/*
		just for debug
	 */
	fmt.Println("EVENTLOOP STOPPED")
}