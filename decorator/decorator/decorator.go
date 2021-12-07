package main

import "fmt"

type Shape interface {
	Render() string
}

type Circle struct {
	Radius float32
}

func (c *Circle) Render() string {
	return fmt.Sprintf("Circle of radius: %.2f", c.Radius)
}

func (c *Circle) Resize(factor float32) {
	c.Radius *= factor
}

type Square struct {
	Side float32
}

func (s *Square) Render() string {
	return fmt.Sprintf("Square with side: %.2f", s.Side)
}

type ColoredShape struct {
	Shape Shape
	Color string
}

func (cs *ColoredShape) Render() string {
	return fmt.Sprintf("%s, which also has the color: %s",
		cs.Shape.Render(), cs.Color)
}

type TransparentShape struct {
	Shape        Shape
	Transparency float32
}

func (ts *TransparentShape) Render() string {
	return fmt.Sprintf("%s, which has %d%% transparency",
		ts.Shape.Render(), int(ts.Transparency * 100.))
}

func main() {
	circle := Circle{2.}
	fmt.Println(circle.Render())

	redCircle := ColoredShape{&circle, "Red"}
	fmt.Println(redCircle.Render())

	rhCircle := TransparentShape{&redCircle, .5}
	fmt.Println(rhCircle.Render())
}
