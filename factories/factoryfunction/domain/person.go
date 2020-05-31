package domain

import (
	"fmt"

	"github.com/isfanazha/go-design-patterns/factories/factoryfunction/usecase"
)

type Person struct {
	Name string
	Age  int
}

// Factory function
// Approach 1 - Using struct type
func NewPerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

// Approach 2 - Using interface
func NewPerson2(name string, age int) usecase.Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person) SayHello() {
	fmt.Printf("Hello I'm %v and I'm %v years old", p.Name, p.Age)
}
