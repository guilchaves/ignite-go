package main

import "fmt"

func main() {
	foo(123)
	foo("")
	foo([]int{1, 2, 3})

	// ~string
	mt := MyType("hello")
	foo(mt)

	ms := MyStruct[int]{Foo: 123}
	ms.foo()
}

type MyType string

type MyConstraint interface {
	~int | ~string | ~[]int
}

func foo[T MyConstraint](arg T) {
	fmt.Println(arg)
}

type MyStruct[T any] struct {
	Foo T
}

func (MyStruct[T]) foo() {}

func Contains[T comparable](s []T, cmp T) bool {
	for _, str := range s {
		if str == cmp {
			return true
		}
	}
	return false
}
