package main

import "fmt"

func main() {
	/*
    var res int

	One way of writing loops (like C or js)
	   for i := 0; i < 10; i++ {
	       res++
	   }

    Another way
        var i int
	    for ; i < 10; i++ {
	    	       res++
	    	   }

    This works as a while true
        for {
            res++
            i++
        }

    This also works
    for i < 10 {
    		res++
    		i++
    	}
	*/

    arr := [10]int{1,2,3,4,5,6,7,8,9,10}
    for range arr {
        fmt.Println("inside loop")
    }

    // print indexes and elements
    for i, elem := range arr {
        fmt.Println(i, elem)
    }

    // print elements only (_ is a blank identifier)
    for _, elem := range arr {
        fmt.Println(elem)
    }
}
