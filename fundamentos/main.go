package main

import (
	"fmt"
	"fundamentos/functions"
	"fundamentos/types"
	"strconv"
)

// package variables must be defined as so, never using short syntax
var age int = 30

func main() {
    printFunctions()
    printVariables()
    printTypes()
    printConvertedTypes()
    }

func printFunctions(){
    fmt.Println("==========FUNCTIONS=========")
    fmt.Println("split: ")
    res, rem := functions.Split(10,3)
    fmt.Println(res, rem)
    fmt.Println("sum: ")
    x := functions.Sum(2)(1)
    fmt.Println(x)
    fmt.Println("sumB: ")
    fmt.Println(functions.SumB(10, 10, age))
}

func printVariables() {
    fmt.Println("==========VARIABLES=========")
    foo, bar := "foo", 40
    fmt.Println(foo, bar)
}

func printTypes(){
    fmt.Println("==========TYPES=========")
    var b uint8 = 10
    types.TakeByte(b)

    var x int = 65
    f := float64(x)
    fmt.Println(f)

    var y int = 10084 
    s := string(y) // this will print the unicode character for the number 10084 
    s2 := strconv.FormatInt(int64(y), 10)

    fmt.Println(s)
    fmt.Println(s2)
}

func printConvertedTypes(){
    fmt.Println("==========CONVERTED-TYPES=========")
    const x = 3.14
    types.TakeInt32(10)
    types.TakeInt64(10)
    types.TakeString("foo")
    types.TakeFloat64(x)
}






