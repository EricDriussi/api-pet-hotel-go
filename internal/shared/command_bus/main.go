package commandbus

import "context"

type Handler interface {
	Handle(context.Context, Command) error
}
type Type string

type Command interface {
	Type() Type
}

type CommandBus interface {
	Dispatch(context.Context, Command) error
	Register(Type, Handler)
}
