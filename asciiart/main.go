package main

import (
	"fmt"
)

func main() {
	for i := 1; i <= 10; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	fmt.Println()

	for i := 10; i >= 1; i-- {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	fmt.Println()

	for i := 10; i >= 1; i-- {
		for j := 0; j < 10-i; j++ {
			fmt.Print(" ")
		}

		for k := 0; k < i; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}

	fmt.Println()

	for i := 10; i >= 1; i-- {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}

		for k := 0; k < 10-i; k++ {
			fmt.Print("*")
		}

		fmt.Println()
	}
}
