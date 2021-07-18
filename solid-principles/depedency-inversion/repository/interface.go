package repository

import "github.com/isfanazha/go-design-patterns/solid-principles/depedency-inversion/domain"

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*domain.Person
}
