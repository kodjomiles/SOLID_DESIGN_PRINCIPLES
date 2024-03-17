package main

import (
	"fmt"
)

//Open Close Principle
//Open for extension closed for modification
//The concept of open close is well illustrated using Enterprise Specification Pattern

// Color type
type Color int

const (
	red Color = iota
	green
	blue
)

// Size type
type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

//tightly coupled, and we will break the open close on this type if we keep adding other filter criteria like filterbysize or by size and color

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	results := make([]*Product, 0)

	for _, product := range products {
		if product.color == color {
			results = append(results, &product)
		}
	}

	return results
}

//breaks the open close

func (f *Filter) FilterBySize(size Size, products []Product) []*Product {
	results := make([]*Product, 0)
	for _, product := range products {
		if product.size == size {
			results = append(results, &product)
		}
	}
	return results
}

// instead, what we can do is to create a specification interface that checks if the product isSatified to use that interface

type Specification interface {
	isSatisfied(product *Product) bool
}

// and then we can have color specification

type ColorSpecification struct {
	color Color
}

func NewColorSpecification(color Color) *ColorSpecification {
	return &ColorSpecification{color: color}
}

func (c ColorSpecification) isSatisfied(product *Product) bool {
	return c.color == product.color
}

type SizeSpecification struct {
	size Size
}

func NewSizeSpecification(size Size) *SizeSpecification {
	return &SizeSpecification{size: size}
}

func (s SizeSpecification) isSatisfied(product *Product) bool {
	return s.size == product.size
}

//doing this is not the best, but it will be best to use a composite specification
// like an and specification and an or specification

type ColorSizeSpecification struct {
	color Color
	size  Size
}

func NewColorSizeSpecification(color Color, size Size) *ColorSizeSpecification {
	return &ColorSizeSpecification{color: color, size: size}
}

func (c ColorSizeSpecification) isSatisfied(product *Product) bool {
	return c.color == product.color && c.size == product.size
}

//composite specification

type AndSpecification struct {
	first, second Specification
}

func (a AndSpecification) isSatisfied(product *Product) bool {
	return a.first.isSatisfied(product) && a.second.isSatisfied(product)
}

// BetterFilter Better filter
type BetterFilter struct {
}

func (b *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	results := make([]*Product, 0)
	for _, product := range products {
		if spec.isSatisfied(&product) {
			results = append(results, &product)
		}
	}
	return results
}

func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, medium}
	melon := Product{"Melon", green, large}
	house := Product{"House", blue, large}
	car := Product{"Car", red, medium}

	products := []Product{apple, tree, house, melon, car}

	f := Filter{}
	bf := BetterFilter{}

	fmt.Printf("Green product: (OLD):\n")
	for _, product := range f.FilterByColor(products, green) {
		fmt.Printf("-%s is green\n", product.name)
	}

	greenColorSpec := NewColorSpecification(green)
	mediumSizeSpec := NewSizeSpecification(medium)
	greenLargeSizeSpec := NewColorSizeSpecification(green, large)
	mediumGreenSpec := AndSpecification{greenColorSpec, mediumSizeSpec}

	fmt.Printf("Green product: (New):\n")
	for _, product := range bf.Filter(products, greenColorSpec) {
		fmt.Printf("-%s is green\n", product.name)
	}

	fmt.Printf("Medium product: (New):\n")
	for _, product := range bf.Filter(products, mediumSizeSpec) {
		fmt.Printf("-%s is medium\n", product.name)
	}

	fmt.Printf("Green large product: (not too ideal):\n")
	for _, product := range bf.Filter(products, greenLargeSizeSpec) {
		fmt.Printf("-%s is green and large\n", product.name)
	}

	fmt.Printf("Green medium product: (bestway):\n")
	for _, product := range bf.Filter(products, mediumGreenSpec) {
		fmt.Printf("-%s is green and medium\n", product.name)
	}

}
