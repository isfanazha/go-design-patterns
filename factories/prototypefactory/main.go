package main

import (
	"fmt"

	"github.com/isfanazha/go-design-patterns/factories/prototypefactory/domain"
)

func main() {
	m := domain.NewEmployee(domain.Manager)
	m.Name = "Sam"
	fmt.Println(m)

	e := domain.NewEmployee(domain.Developer)
	e.Name = "John"
	fmt.Println(e)
}
