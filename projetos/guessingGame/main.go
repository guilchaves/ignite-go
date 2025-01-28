package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guessing game")
	fmt.Println("A random number will be generated between 0 and 100. Try to guess it.")

	randomNum := rand.Int64N(101)
	scanner := bufio.NewScanner(os.Stdin)
	guesses := [10]int64{}

	for i := range guesses {
		fmt.Println("Enter your guess from 0 to 100: ")
		scanner.Scan()
		guess := scanner.Text()
		guess = strings.TrimSpace(guess)
		guessInt, err := strconv.ParseInt(guess, 10, 64)

		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			return
		}

		switch {
		case guessInt < randomNum:
			fmt.Println("Wrong. Too low.")
		case guessInt > randomNum:
			fmt.Println("Wrong. Too high.")
		case guessInt == randomNum:
			fmt.Printf(
				"Correct! The number was: %d\n"+
					"You guessed it in %d tries.\n"+
					"Your guesses: %v\n",
				randomNum, i+1, guesses[:i],
			)
			return
		}

		guesses[i] = guessInt
	}

	fmt.Printf(
		"You have reached the maximum number of guesses. The correct number was %d.\n"+
			"Your guesses: %v\n",
		randomNum, guesses,
	)
}
