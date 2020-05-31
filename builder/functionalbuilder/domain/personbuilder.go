package domain

type PersonMod func(*Person)

type PersonBuilder struct {
	Actions []PersonMod
}

func (pb *PersonBuilder) Called(name string) *PersonBuilder {
	pb.Actions = append(pb.Actions, func(person *Person) {
		person.Name = name
	})

	return pb
}

func (pb *PersonBuilder) Build() *Person {
	p := Person{}
	for _, a := range pb.Actions {
		a(&p)
	}

	return &p
}

// extend PersonBuilder
func (pb *PersonBuilder) WorksAsA(position string) *PersonBuilder {
	pb.Actions = append(pb.Actions, func(person *Person) {
		person.Position = position
	})

	return pb
}
