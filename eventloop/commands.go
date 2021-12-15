package eventloop

type Command interface {
	Execute(h Handler)
}

type Handler interface {
	Post(command Command, isInner bool)
}

type CommandFunc func(h Handler)

func (cf CommandFunc) Execute(h Handler) {
	cf(h)
}