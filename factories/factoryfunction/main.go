package main

import (
	"fmt"

	"github.com/isfanazha/go-design-patterns/factories/factoryfunction/domain"
)

func main() {
	// Initialize directly
	p := domain.Person{Name: "Cristiano", Age: 22}
	fmt.Println(p)

	// Use a constructor - return struct type
	p2 := domain.NewPerson("Ronaldo", 21)
	p2.Age = 30
	fmt.Println(p2)

	// Use a constructor - return interface
	p3 := domain.NewPerson2("Lionel", 20)
	p3.SayHello()
}
