package commands

import "github.com/EricDriussi/api-pet-hotel-go/internal/shared/command_bus"

const CreateBookingCommandType commandbus.Type = "command.create.booking"

type CreateBookingCommand struct {
	Id       string
	PetName  string
	Duration string
}

func NewCreateBooking(id, name, duration string) CreateBookingCommand {
	return CreateBookingCommand{
		Id:       id,
		PetName:  name,
		Duration: duration,
	}
}

func (c CreateBookingCommand) Type() commandbus.Type {
	return CreateBookingCommandType
}
