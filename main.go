package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) string {
	fmt.Print(prompt)
	input, _ := r.ReadString('\n')
	return strings.TrimSpace(input)
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill:", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Options:")
	fmt.Println("\t1. Add item")
	fmt.Println("\t2. Update tip")
	fmt.Println("\t3. Exit")

	input := getInput("Enter your choice: ", reader)

	switch input {
	case "1":
		// addItem(b)
	case "2":
		// updateTip(b)
	case "3":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid input")
		promptOptions(b)
	}
}

func main() {

	myBill := createBill()
	promptOptions(myBill)

}
