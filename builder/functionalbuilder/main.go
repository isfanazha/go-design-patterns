package main

import (
	"fmt"

	"github.com/isfanazha/go-design-patterns/builder/functionalbuilder/domain"
)

func main() {
	b := domain.PersonBuilder{}
	p := b.Called("Cristiano").WorksAsA("football player").Build()
	fmt.Println(*p)
}
