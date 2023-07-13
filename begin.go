package begin

import "github.com/google/uuid"

const accessAge uint = 18

func New(name string, age uint) *Profile {
	return &Profile{
		Uuid: uuid.New(),
		Name: name,
		Age:  age,
	}
}

type Profile struct {
	Uuid uuid.UUID
	Name string
	Age  uint
}

func (p Profile) GetName() string {
	return p.Name
}

func (p Profile) GetUuid() string {
	return p.Uuid.String()
}

func (p Profile) GetAge() uint {
	return p.Age
}

func (p Profile) AcceptAge() bool {
	return p.GetAge() > accessAge
}
