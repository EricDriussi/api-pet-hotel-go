package commandbus

import "context"

type Handler interface {
	Handle(context.Context, Command) error
}
