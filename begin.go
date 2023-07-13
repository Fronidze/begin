package begin

import "github.com/google/uuid"

type Profile struct {
	Uuid uuid.UUID
	Name string
	Age  uint
}

func (profile *Profile) New(name string, age uint) *Profile {
	return &Profile{
		Uuid: uuid.New(),
		Name: name,
		Age:  age,
	}
}
