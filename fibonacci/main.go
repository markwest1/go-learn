package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	ul, err := getUserInput()

	for err != nil {
		fmt.Print("Invalid input: ")
		fmt.Println(err)
		ul, err = getUserInput()
	}

	fmt.Println(ul)
}

func getUserInput() (n int, err error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Print fibonacci sequence up to: ")
	scanner.Scan()

	text := scanner.Text()
	if err := scanner.Err(); err != nil {
		return 0, err
	}

	return strconv.Atoi(text)
}
