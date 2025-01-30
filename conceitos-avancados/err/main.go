/* Primeira parte da aula: Entendendo errors

package main

import (
	"errors"
	"fmt"
)

func main() {
	user, err := NewUser(true)

	if err != nil {
		fmt.Println("An error occurred")
		return
	}

	user.Foo()
}

type User struct {
	foo string
}

func (u User) Foo() {
	fmt.Println(u.foo)
}

func NewUser(wantErr bool) (*User, error) {
	if wantErr {
		return nil, errors.New("error")
	}

	return &User{}, nil
}



Segunda parte da aula: Criando erros

package main

import (
	"errors"
	"math"
)

type SqrtError struct {
	msg string
}

func (s SqrtError) Error() string { return s.msg }

func squareRoot(x float64) (float64, error) {
	if x < 0 {
		return 0, SqrtError{"there is no square root of a negative number"}
	}
	result := math.Sqrt(x)
	return result, nil
}

var ErrNotFound = errors.New("not found")

func main() {
	//	x := -10
	//	res, err := squareRoot(float64(x))
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	fmt.Println(res)

	//errors.Is example:
		    err := foo()
			if err != nil && errors.Is(err, ErrNotFound) {
				fmt.Println("Error not found!")
			}
			fmt.Println("Not error on foo")

	//errors.As example
			err := foo()
			var sqrtError SqrtError
			if err != nil && errors.As(err, &sqrtError) {
				fmt.Println(sqrtError.msg)
				return
			}

		    fmt.Println("Not an error on foo")



//func foo() error { return SqrtError{msg: "test"} }*/

// Terceira parte da aula - error wrapping
package main

import (
	"errors"
	"fmt"
)

func main() {
	err := foo()
	if err != nil {
		fmt.Println("error occured: ", err)
		return
	}

}

var ErrFoo = errors.New("🦆 quack quack")

func foo() error {
	err := bar()
	if err != nil {
		return fmt.Errorf("an error occured in foo: %w", err)
	}
	return nil
}

func bar() error {
	return ErrFoo
}
