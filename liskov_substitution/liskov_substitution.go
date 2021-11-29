package main

import (
	"fmt"
)

// Liskov Substitution Principle

type Sized interface {
	GetWidth() int
	SetWidth(int)
	GetHeight() int
	SetHeight(int)
}

type Rectangle struct {
	width int
	height int
}

func (r *Rectangle) GetWidth() int {
	return r.width
}

func (r *Rectangle) SetWidth(width int) {
	r.width = r.width
}

func (r *Rectangle) GetHeight() int {
	return r.height
}

func (r *Rectangle) SetHeight(height int) {
	r.height = height
}

type Square struct {
	Rectangle
}

func NewSquare(size int) *Square{
	square := Square{}
	square.width = size
	square.height = square.width

	return &square
}

func (square *Square) SetWidth(width int) {
	square.width = width
	square.height = width
}

func (square *Square) SetHeight(height int) {
	square.width = height
	square.height = height
}

func UseIt(sized Sized) {
	width := sized.GetWidth()
	sized.SetHeight(10)
	expectedArea := 10 * width
	actualArea := sized.GetWidth() * sized.GetHeight()
	fmt.Print("Expected an area of ", expectedArea, 
	" but got ", actualArea, "\n")
}

type Square2 struct {
	size int
}

func (s *Square2) Rectangle() Rectangle {
	return Rectangle{s.size, s.size}
}

func main() {
	rectangle := &Rectangle{2, 3}
	UseIt(rectangle)

	square := NewSquare(2)
	UseIt(square)
}
