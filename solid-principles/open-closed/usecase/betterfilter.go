package usecase

import "github.com/isfanazha/go-design-patterns/solid-principles/open-closed/domain"

// In BetterFilter we use Specification interface as parameter which is more flexible,
// So the implementation depends on the type (ColorSpecification, SizeSpecification, AndSpecification).
// If you want to filter by particular new type, just create a new Specification.

type BetterFilter struct{}

func (f *BetterFilter) Filter(products []domain.Product, spec Specification) []*domain.Product {
	result := make([]*domain.Product, 0)

	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}

	return result
}
