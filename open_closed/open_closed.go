package main

import "fmt"

// OCP
// Open for extension, closed for modification
// Specification

type Colour int

const (
	red Colour = iota
	green
	blue
)

type Size int

const (
	small Size = iota
	medium
	large
)

type Product struct {
	name string
	colour Colour
	size Size
}

type Filter struct {

}

func (f *Filter) FilterByColor(products []Product, colour Colour) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if v.colour == colour {
			result = append(result, &products[i])
		}
	}

	return result
}

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColourSpecification struct {
	colour Colour
}

func (cs ColourSpecification) IsSatisfied(p *Product) bool {
	return p.colour == cs.colour
}

type SizeSpecification struct {
	size Size
}

func (ss SizeSpecification) IsSatisfied(p *Product) bool {
	return p.size == ss.size
}

type BetterFilter struct {}

func (f *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	result := make([]*Product, 0)

	for i, v := range products {
		if spec.IsSatisfied(&v) {
			result = append(result, &products[i])
		}
	}

	return result
}

type AndSpecification struct {
	specs []Specification
}

func (asp AndSpecification) IsSatisfied(p *Product) bool {
	result := true

	for _, spec := range asp.specs {
		if (!spec.IsSatisfied(p)) {
			result = false
			break
		}
	}

	return result
}


func main() {
	apple := Product{"Apple", green, small}
	tree := Product{"Tree", green, large}
	house := Product{"House", blue, large}

	products := []Product{apple, tree, house}
	fmt.Printf("Green products (old):\n")
	f := Filter{}
	for _, v := range f.FilterByColor(products, green) {
		fmt.Printf("- %s is green\n", v.name)
	}

	fmt.Printf("Green products (new):\n")
	greenSpec := ColourSpecification{green}
	bf := BetterFilter{}
	for _, v := range bf.Filter(products, greenSpec) {
		fmt.Printf("- %s is green\n", v.name)
	}

	fmt.Printf("Green and large products (new):\n")
	sizeSpec := SizeSpecification{large}
	andSpec := AndSpecification{[]Specification{greenSpec, sizeSpec}} 
	bf2 := BetterFilter{}
	for _, v := range bf2.Filter(products, andSpec) {
		fmt.Printf("- %s is green and large\n", v.name)
	}
}
