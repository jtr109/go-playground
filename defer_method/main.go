package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Hello() {
	fmt.Println("Hello, my name is", p.Name)
}

func NewPerson(name string) *Person {
	return &Person{name}
}

func (p *Person) Rename(name string) {
	p.Name = name
}

func main() {
	fmt.Println("Hello world!")
	p := NewPerson("Tom")
	defer p.Hello()
	p.Rename("Jerry")
}
