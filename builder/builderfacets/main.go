package main

import (
	"fmt"

	"github.com/isfanazha/go-design-patterns/builder/builderfacets/domain"
)

func main() {

	pb := domain.NewPersonBuilder()

	// Builder facets
	pb.
		Lives().
		At("Jl. Kebenaran").
		In("Jakarta").
		WithPostcode("12345").
		Works().
		At("Factory").
		AsA("Programmer").
		Earning(1000)

	person := pb.Build()

	fmt.Println(*person)
}
