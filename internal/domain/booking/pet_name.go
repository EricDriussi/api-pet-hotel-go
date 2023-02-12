package domain

type petName struct {
	value string
}

func newPetName(value string) (petName, error) {
	if value == "" {
		return petName{}, ErrEmptyPetName
	}

	return petName{
		value: value,
	}, nil
}

func (name petName) string() string {
	return name.value
}
