package main

import "fmt"

func main() {
    var z *int

    x := 10
    take(&x)

    y := create()

    // Not allowed, because you are dereferencing a nil pointer
    //output: the program will panic
    take(z)

    fmt.Println(x)
    fmt.Println(*y)
}

func take(x *int) {
    *x = 100
}

func create() *int {
    x := 10
    return &x
}


