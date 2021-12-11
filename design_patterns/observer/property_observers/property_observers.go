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
	Name string
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
	if (age == p.age) {
		return
	}
	p.age = age
	pc := PropertyChange{"Age", p.age}
	p.Fire(pc)
}

type TrafficManagent struct {
	o *Observable
}

func (t *TrafficManagent) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Value.(int) >= 18 {
			fmt.Println("Congrats, you can drive now!")
			t.o.Unsubscribe(t)
		}
	}
}

func main() {
	p := NewPerson(15)
	t := &TrafficManagent{&p.Observable}
	p.Subscribe(t)

	for age := 16; age <= 20; age++ {
		fmt.Println("Setting the age to", age)
		p.SetAge(age)
	}
}
