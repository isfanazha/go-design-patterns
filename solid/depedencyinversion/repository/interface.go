package repository

import "github.com/isfanazha/go-design-patterns/solid/depedencyinversion/domain"

type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*domain.Person
}
