package usecase

import "github.com/isfanazha/go-design-patterns/solid-principles/open-closed/domain"

// Filter has the following functions:
// 1. FilterByColor: Filter product by color
// 2. FilterBySize: Filter product by size
// 3. FilterBySizeAndColor: Filter product by size and color
// We always add new function and logic when there is a new filter.
// Open closed principle is all about being open for extension,
// but in filter struct we always keep add new function when we want to add new filter.
// So we can improve this using specification pattern.

type Filter struct{}

func (f *Filter) FilterByColor(products []domain.Product, color domain.Color) []*domain.Product {
	result := make([]*domain.Product, 0)

	for i, v := range products {
		if v.Color == color {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) FilterBySize(products []domain.Product, size domain.Size) []*domain.Product {
	result := make([]*domain.Product, 0)

	for i, v := range products {
		if v.Size == size {
			result = append(result, &products[i])
		}
	}

	return result
}

func (f *Filter) FilterBySizeAndColor(products []domain.Product, size domain.Size, color domain.Color) []*domain.Product {
	result := make([]*domain.Product, 0)

	for i, v := range products {
		if v.Size == size && v.Color == color {
			result = append(result, &products[i])
		}
	}

	return result
}
