package main

import (
	"fmt"
	"math"
)

func main() {

	// Initilizing a map with map literal
	names := map[string]string{
		"Pedro":   "Pessoa",
		"Joaquim": "Pedro",
	}

	// Initilizing a map with make function
	m := make(map[string]string, 100)

	// Initilizing a map with slices
	slices := map[string][]int{
		"Pedro": {1, 2, 3},
	}

	// Another example of creating and getting value from a map
	myMap := make(map[string]string)
	myMap["Pedro"] = "Pessoa"
	value, ok := myMap["Pedro"]
	value2, ok := myMap["Joaquim"]

	fmt.Println(names)
	fmt.Println(m)
	fmt.Println(slices)
	fmt.Println(value, ok)
	fmt.Println(value2, ok)

	delete(myMap, "Pedro")
	value, ok = myMap["Pedro"]
	fmt.Println(value, ok)

	// Clear the map but keep its capacity
	clear(myMap)
	fmt.Println(myMap)

	fmt.Println("====FLOAT EXAMPLE====")
	var f float64

	f = math.NaN()

	var f2 float64
	f2 = math.NaN()

	fmt.Println(f == f2) // false

    m2 := map[float64]string{
        f: "Pedro",
        f2: "Pessoa",
    }

    fmt.Println(m2)
    value3, ok := m2[f]
    fmt.Println(value3, ok)
    delete(m2, f)
    fmt.Print("After delete: ")
    fmt.Println(m2)
    clear(m2)

    fmt.Print("After clear: ")
    fmt.Println(m2)


    // Another example: iterating over a map
    fmt.Println("====ITERATING OVER A MAP====")
    m3 := map[string]string{
        "Pedro": "Pessoa",
        "Joaquim": "Pedro",
    }

    for k, v := range m3 {
        fmt.Println(k, v) 
    } 

    for k := range m3 {
        if k == "Pedro" {
            delete(m3, k)
        }
    }

    fmt.Println("After deleting pedro: ")
    fmt.Println(m3)

}
