package commandbus

import "context"

type CommandBus interface {
	Dispatch(context.Context, Command) error
	Register(Type, Handler)
}
