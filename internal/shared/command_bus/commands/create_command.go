package commands

import command "github.com/EricDriussi/api-pet-hotel-go/internal/shared/command_bus/definition"

const CreateBookingCommandType command.Type = "command.create.booking"

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

func (c CreateBookingCommand) Type() command.Type {
	return CreateBookingCommandType
}
