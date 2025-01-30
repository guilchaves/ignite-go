package main

import "fmt"

/*
func main(){

    Aula 01

    arr := [5]int{1, 2, 3, 4, 5}
    slice := arr[1:4]

    arr[2] = 15
    slice[0] = 123

    sliceA := []int{1,2,3}

    // using [:] to copy the whole array
    // other ways to copy the whole array are:
    //  - sliceB := arr[0:len(arr)]
    //  - sliceB := arr[:len(arr)]
    sliceB := arr[:]

    fmt.Println(slice) // [2 15 4]
    fmt.Println(arr) // [1 123 15 4 5]
    fmt.Println(sliceA) // [1 2 3]
    fmt.Println(sliceB) // [1 123 15 4 5]

    sliceX := []int{1,2,3,4,5}
    sliceY := sliceX[:0]
    fmt.Println(sliceY, len(sliceY), cap(sliceY)) // [] 0 5

    var nilSlice []int
    fmt.Println(nilSlice == nil) // true
}
*/


/*
Aula 02

var filmsOnDb = []string{"Star Wars", "The Godfather", "The Matrix", "The Lord of the Rings", "The Shawshank Redemption", "Pulp Fiction", "Fight Club", "Inception", "The Dark Knight", "The Silence of the Lambs" }

func main() {
	resultsFromApi := []int{1, 2, 3, 4, 5}

    films := make([]string, 0, 10)

	for _, id := range resultsFromApi {
		film := filmsOnDb[id]
		fmt.Println(len(films), cap(films))
		films = append(films, film)
		fmt.Println(len(films), cap(films))
	}

	fmt.Println(films)
}
*/

// Aula 03
func main(){
    arr := [4]int{1, 2, 3, 4}
    slice := arr[:2:2] // [1 2] 2
    fmt.Println(slice, cap(slice))

    foo(&arr)
    fmt.Println(arr)
}

func foo(slice *[4]int){
    slice[0] = 123
}

