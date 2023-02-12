package domain

import "errors"

var ErrEmptyPetName = errors.New("the field Pet Name can not be empty")

type PetName struct {
	value string
}

func NewPetName(value string) (PetName, error) {
	if value == "" {
		return PetName{}, ErrEmptyPetName
	}

	return PetName{
		value: value,
	}, nil
}

func (name PetName) String() string {
	return name.value
}
