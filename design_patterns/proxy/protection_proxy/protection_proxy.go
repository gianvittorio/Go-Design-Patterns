package main

import "fmt"

type Driven interface {
	Drive()
}

type Car struct {

}

func (c *Car) Drive() {
	fmt.Println("The car is being driven")
}

type Driver struct {
	Age int
}

type CarProxy struct {
	car *Car
	driver *Driver
}

func NewCarProxy(driver *Driver) *CarProxy {
	return &CarProxy{&Car{}, driver}
}

func (cp *CarProxy) Drive() {
	if (cp.driver.Age >= 16) {
		cp.car.Drive()
		return
	}

	fmt.Println("Driver too young!")
}

func main() {
	driver := Driver{22}
	car := NewCarProxy(&driver)
	car.Drive()
}
