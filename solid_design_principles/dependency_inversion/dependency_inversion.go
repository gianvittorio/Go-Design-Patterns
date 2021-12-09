package main

import "fmt"

// Dependency Inversion Principle
// HLM should not depend on LLM
// Both should depend on abstractions

type Relationship int

const (
	Parent Relationship = iota
	Child
	Silbling
)

type Person struct {
	name string
	//
}

type Info struct {
	from         *Person
	relationship Relationship
	to           *Person
}

// low-level module
type RelationshipBrowser interface {
	FindAllChildrenOf(name string) []*Person
}

type Relationships struct {
	relationships []Info
}

func (r *Relationships) FindAllChildrenOf(name string) []*Person {
	result := make([]*Person, 0)

	for i, v := range r.relationships {
		if v.relationship == Parent &&
		v.from.name == name {
			result = append(result, r.relationships[i].to)
		}
	}

	return result
}

func (r *Relationships) AddParentAndChild(parent, child *Person) {
	r.relationships = append(
		r.relationships,
		Info{parent, Parent, child})
	r.relationships = append(
		r.relationships,
		Info{child, Child, parent})
}

// high-level module
type Research struct {
	// break DIP
	//relationships Relationships
	browser RelationshipBrowser
}

func (r *Research) Investigate() {
	for _, p := range r.browser.FindAllChildrenOf("John") {
		fmt.Println("John has a child called: ", p.name)
	}
}

func main() {
	parent := Person{"John"}
	child1 := Person{"Chris"}
	child2 := Person{"Matt"}

	relationships := Relationships{}
	relationships.AddParentAndChild(&parent, &child1)
	relationships.AddParentAndChild(&parent, &child2)

	r := Research{&relationships}
	r.Investigate()
}
