package eventloop

import (
	"sync"

	engine "github.com/UniversalCorn/lab4/commands"
)

type Queue struct {
	sync.Mutex

	commands []engine.Command
}

func (q *Queue) push(command engine.Command) {
	q.Lock()
	defer q.Unlock()

	q.commands = append(q.commands, command)
}

func (q *Queue) pop() engine.Command {
	q.Lock()
	defer q.Unlock()

	command := q.commands[0]
	q.commands[0] = nil
	q.commands = q.commands[1:]
	return command
}

func (q *Queue) empty() bool {
	q.Lock()
	defer q.Unlock()

	return len(q.commands) == 0
}
