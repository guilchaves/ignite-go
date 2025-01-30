package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	start := time.Now()
    const n = 10
    var wg sync.WaitGroup
    wg.Add(n)

	for range n {
		go func() {
            defer wg.Done()

			res, err := http.Get("https://google.com")
			if err != nil {
				panic(err)
			}

			defer res.Body.Close()
			fmt.Println("ok")
		}()
	}

    wg.Wait()
	fmt.Println("Time taken: ", time.Since(start))
}
