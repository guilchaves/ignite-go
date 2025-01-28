package main

import (
	"errors"
	"fmt"
	"math"
)

func main(){
    if x := math.Sqrt(4); x < 1 {
        fmt.Println(x)
    } else if x > 1 {
        fmt.Println("x is greater than 0")
    } else {
        fmt.Println("x is less than 0")
    }

    if err := doError(); err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("no error")
    }

    myBool := true
    if myBool {
        fmt.Println("myBool is true")
    }
    
}

func doError() error {
    return errors.New("error was thrown")
}
