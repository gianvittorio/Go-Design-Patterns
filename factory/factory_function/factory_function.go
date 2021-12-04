package main

type Person struct {
	Name string
	Age int
	EyeCount int
}

func NewPerson(name string, age int) *Person {
	if age < 16 {
		panic("person is underraged!")
	}
	
	return &Person{name, age, 2}
}

func main() {
	p := NewPerson("John", 22)
}
