package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	fmt.Println(person{"Bob", 42})
	fmt.Println(person{name: "Alfred", age: 24})
	fmt.Println(&person{name: "Ann", age: 40})
	fmt.Println(newPerson("Alfred", 30))

	s := person{name: "Sean", age: 40}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)

	dog := struct {
		name   string
		isGood bool
	}{
		"Rex",
		true,
	}
	fmt.Println(dog)
}
