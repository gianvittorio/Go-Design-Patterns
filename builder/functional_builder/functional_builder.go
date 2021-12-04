package main

import "fmt"

type Person struct {
	name, position string
}

type personMod func(*Person)

type PersonBuilder struct {
	actions []personMod
}

func (b *PersonBuilder) Called(name string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.name = name
	})

	return b
}

func (b *PersonBuilder) WorksAsA(position string) *PersonBuilder {
	b.actions = append(b.actions, func(p *Person) {
		p.position = position
	})

	return b
}

func (b *PersonBuilder) Build() *Person {
	person := new(Person)

	for _, action := range b.actions {
		action(person)
	}

	return person
}

func main() {
	b := PersonBuilder{}
	p := b.Called("Dmitri").
		WorksAsA("Programmer").
		Build()
	fmt.Println(*p)
}
