package main

import (
	"fmt"

	"github.com/isfanazha/go-design-patterns/prototype/prototypefactory/domain"
)

func main() {
	john := domain.NewMainOfficeEmployee("John", 100)
	jane := domain.NewAuxOfficeEmployee("Jane", 200)

	fmt.Println(john)
	fmt.Println(jane)
}
