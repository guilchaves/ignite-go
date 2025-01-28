package main

import (
	"fmt"
	"time"
)

func main() {
	do(1)
	fmt.Println(isWeekend(time.Now()))
	fmt.Println(isWeekday(time.Now()))
}

func do(x int) {
	switch x {
	case 1:
		fmt.Println("one")
		fallthrough
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("many")
	}
}

func isWeekend(day time.Time) bool {
	switch {
	case day.Weekday() > 0 && day.Weekday() < 6:
		return false
	default:
		return true
	}
}

func isWeekday(day time.Time) bool {
	switch day.Weekday() {
	case time.Saturday, time.Sunday:
		return false
	default:
		return true
	}
}

func findType(x any) {
	switch t := x.(type) {
	case string:
	case int:
	case nil:
	}
}
