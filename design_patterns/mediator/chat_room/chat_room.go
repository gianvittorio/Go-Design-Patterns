package main

import "fmt"

type Person struct {
	Name    string
	Room    *ChatRoom
	chatLog []string
}

func NewPerson(name string) *Person {
	return &Person{Name: name}
}

func (p *Person) Receive(sender, message string) {
	s := fmt.Sprintf("%s: %s", sender, message)
	fmt.Printf("[%s's chat session]: %s\n", p.Name, s)
	p.chatLog = append(p.chatLog, s)
}

func (p *Person) Say(message string) {
	p.Room.Broadcast(p.Name, message)
}

func (p *Person) PrivateMessage(who, message string) {
	p.Room.PrivateMessage(p.Name, who, message)
}

type ChatRoom struct {
	people []*Person
}

func (c *ChatRoom) Broadcast(from, message string) {
	for _, p := range c.people {
		if p.Name == from {
			continue
		}
		p.Receive(from, message)
	}
}

func (c *ChatRoom) PrivateMessage(from, to, message string) {
	for pos := 0; pos < len(c.people); pos++ {
		person := c.people[pos]
		if (person.Name == to) {
			person.Receive(from, message)
		}
	}
}

func (c *ChatRoom) Join(p *Person) {
	joinMsg := p.Name + " joins the chat"
	c.Broadcast("Room", joinMsg)

	p.Room = c
	c.people = append(c.people, p)
}

func main() {
	room := ChatRoom{}

	john := NewPerson("John")
	jane := NewPerson("Jane")

	room.Join(john)
	room.Join(jane)

	john.Say("Hi room")
	jane.Say("Oh, hey John")

	simon := NewPerson("Simon")
	room.Join(simon)
	simon.Say("Hi everyone!")

	jane.PrivateMessage(jane.Name, "Simon, glad you could join us!")
}
