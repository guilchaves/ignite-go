package main

import (
	"fmt"
)

func main() {
    x := doDeferWithReturn()
    fmt.Println(x)

    doDefer()
    doDeferLambda()
    doDeferClojure()
}

// defer works as a stack (LIFO order), last in first out
//output: 123
func doDefer(){
    defer fmt.Println(3)
    defer fmt.Println(2)
    fmt.Println(1)
}

func doDeferWithReturn() int{
	defer fmt.Println("world")
	fmt.Println("hello")
    return 10
}

// params are evaluated at the time of defer statement
// output: 50 10
func doDeferLambda(){
    x := 10
    defer func(y int){
        fmt.Println(y)
    }(x)

    x = 50
    fmt.Println(x)
}

// params are evaluated at the time of defer statement, since x is passed by reference, it will print the latest value of x
// output: 50 50
func doDeferClojure(){
    x := 10
    defer func(){
        fmt.Println(x)
    }()

    x = 50
    fmt.Println(x)
}

