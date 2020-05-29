package main

import (
	"fmt"

	"github.com/isfanazha/go-design-patterns/solid/openclosed/domain"
	"github.com/isfanazha/go-design-patterns/solid/openclosed/usecase"
)

func main() {
	apple := domain.Product{Name: "Apple", Color: domain.Red, Size: domain.Small}
	tree := domain.Product{Name: "Tree", Color: domain.Green, Size: domain.Small}
	house := domain.Product{Name: "House", Color: domain.Green, Size: domain.Large}

	products := []domain.Product{apple, tree, house}

	// Using Filter{}
	fmt.Print("Green products using Filter:\n")
	f := usecase.Filter{}
	for _, v := range f.FilterByColor(products, domain.Green) {
		fmt.Printf(" - %s is green\n", v.Name)
	}

	// Using BetterFilter{}
	fmt.Print("\nGreen products using BetterFilter:\n")
	greenSpec := usecase.ColorSpecification{Color: domain.Green}
	bf := usecase.BetterFilter{}
	for _, v := range bf.Filter(products, &greenSpec) {
		fmt.Printf(" - %s is green\n", v.Name)
	}

	largeSpec := usecase.SizeSpecification{Size: domain.Large}

	largeGreenSpec := usecase.AndSpecification{First: &largeSpec, Second: &greenSpec}
	fmt.Print("\nLarge and blue items using BetterFilter:\n")
	for _, v := range bf.Filter(products, &largeGreenSpec) {
		fmt.Printf(" - %s is large and green\n", v.Name)
	}
}
