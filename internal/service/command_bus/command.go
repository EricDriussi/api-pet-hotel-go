package commandbus

type Type string

type Command interface {
	Type() Type
}
