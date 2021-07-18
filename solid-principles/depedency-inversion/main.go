package main

import (
	"github.com/isfanazha/go-design-patterns/solid-principles/depedency-inversion/domain"
	"github.com/isfanazha/go-design-patterns/solid-principles/depedency-inversion/repository"
	"github.com/isfanazha/go-design-patterns/solid-principles/depedency-inversion/usecase"
)

func main() {
	parent := domain.Person{Name: "John"}
	child1 := domain.Person{Name: "Chris"}
	child2 := domain.Person{Name: "Matt"}

	// Repository is low-level module, ex: database
	relationships := repository.Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	research := usecase.Research{Browser: &relationships}
	research.Investigate()
}
