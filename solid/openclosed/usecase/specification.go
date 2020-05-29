package usecase

import "github.com/isfanazha/go-design-patterns/solid/openclosed/domain"

// Specification interface follows the open closed principles,
// which opens for extension (grab an interface and implement the method somewhere in your code)
// and close to a modification which means once we have designed and test API of particular type,
// we shouldn't really keep jumping into the method and modifying it again again and again.

type Specification interface {
	IsSatisfied(p *domain.Product) bool
}

type ColorSpecification struct {
	Color domain.Color
}

func (spec *ColorSpecification) IsSatisfied(p *domain.Product) bool {
	return p.Color == spec.Color
}

type SizeSpecification struct {
	Size domain.Size
}

func (spec *SizeSpecification) IsSatisfied(p *domain.Product) bool {
	return p.Size == spec.Size
}

type AndSpecification struct {
	First, Second Specification
}

func (spec *AndSpecification) IsSatisfied(p *domain.Product) bool {
	return spec.First.IsSatisfied(p) && spec.Second.IsSatisfied(p)
}
