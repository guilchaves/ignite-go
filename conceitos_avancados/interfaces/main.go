package main

import "fmt"

type Animal interface {
	Sound() string
}

type Dog struct{}
type Cat struct{}

func (Dog) Sound() string {
	return "Woof"
}

func (d *Cat) Sound() string {
	return "Meow"
}

func makeSound(a Animal) {
	fmt.Println(a.Sound())
}

func takeAnimal(a Animal) {
	switch t := a.(type) {
	case *Dog:
        t.Sound()
	case *Cat:
	}

}

func main() {
	//	dog := Dog{}
	//    makeSound(dog)

	var a Animal
	var dog *Dog
	a = dog
	fmt.Println(a)
}
