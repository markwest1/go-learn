package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	ul, err := getUserInput()

	for err != nil {
		if err != nil {
			fmt.Print("Invalid input: ")
			fmt.Println(err)
		}

		ul, err = getUserInput()
	}

	printSequence(ul)
}

func printSequence(upperLimit int) {
	fmt.Print("0, ")

	a := 0
	b := 1

	for b <= upperLimit {
		fmt.Printf("%d", b)
		x := a
		a = b
		b = x + b

		if b <= upperLimit {
			fmt.Print(", ")
		}
	}

	fmt.Println("")
}

func getUserInput() (n int, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Print fibonacci sequence up to: ")
	scanner.Scan()

	text := scanner.Text()
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	n, err = strconv.Atoi(text)

	if err != nil {
		return 0, err
	} else if n <= 0 {
		return n, errors.New("must be greater than or equal to one")
	}

	return n, err
}
