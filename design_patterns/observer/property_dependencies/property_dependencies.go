package main

import (
	"container/list"
	"fmt"
)

type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Observer interface {
	Notify(data interface{})
}

type PropertyChange struct {
	Name  string
	Value interface{}
}

type Person struct {
	Observable
	age int
}

func NewPerson(age int) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		age:        age,
	}
}

func (p *Person) Age() int {
	return p.age
}

func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}

	oldCanVote := p.CanVote()

	p.age = age
	pc := PropertyChange{"Age", p.age}
	p.Fire(pc)

	if oldCanVote != p.CanVote() {
		p.Fire(PropertyChange{"CanVote", p.CanVote()})
	}
}

func (p *Person) CanVote() bool {
	return p.age >= 18
}

type ElectoralRoll struct {
}

func (e *ElectoralRoll) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congratulations, you can vote!")
		}
	}
}

func main() {
	p := NewPerson(0)
	er := &ElectoralRoll{}
	p.Subscribe(er)

	for age := 10; age < 20; age++ {
		fmt.Println("Setting age to", age)
		p.SetAge(age)
	}
}
