package main

import (
	"fmt"
	"bytes"
	"encoding/gob"
)

type DeepCopyable interface {
	DeepCopy() DeepCopyable
}

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)

	_ = e.Encode(p)

	fmt.Println(string(b.Bytes()))

	d := gob.NewDecoder(&b)
	result := new(Person)
	_ = d.Decode(result)
	return result
}

func main() {
	john := Person{"John",
		&Address{"123 London Road", "London", "UK"}, 
		[]string{"Chris", "Matt"}}

	jane := john.DeepCopy()
	jane.Name = "Jane"
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Angela")

	fmt.Println(john, john.Address)
	fmt.Println(jane, jane.Address)
}
