package repository

import "github.com/isfanazha/go-design-patterns/solid-principles/depedency-inversion/domain"

type Relationships struct {
	Relations []domain.Info
}

func (rs *Relationships) AddParentAndChild(parent, child *domain.Person) {
	rs.Relations = append(rs.Relations, domain.Info{From: parent, Relationship: domain.Parent, To: child})
	rs.Relations = append(rs.Relations, domain.Info{From: child, Relationship: domain.Child, To: parent})
}

func (rs *Relationships) FindAllChildrenOf(name string) []*domain.Person {
	result := make([]*domain.Person, 0)

	for i, v := range rs.Relations {
		if v.Relationship == domain.Parent && v.From.Name == name {
			result = append(result, rs.Relations[i].To)
		}
	}

	return result
}
