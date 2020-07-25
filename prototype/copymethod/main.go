package main

import (
	"fmt"

	"github.com/isfanazha/go-design-patterns/prototype/copymethod/domain"
)

func main() {
	john := domain.Person{
		Name:    "John",
		Address: &domain.Address{StreetAddress: "ABC Street", City: "JAKARTA", Country: "ID"},
		Friends: []string{"Chris", "Matt"},
	}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "ZZZ Street"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
