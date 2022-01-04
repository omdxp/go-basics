package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func addItem(b *bill) {
	reader := bufio.NewReader(os.Stdin)

	item := getInput("Enter item name: ", reader)
	price := getInput("Enter item price: ", reader)

	// convert price to float
	priceFloat, err := strconv.ParseFloat(price, 64)

	if err != nil {
		fmt.Println("Invalid price")
		return
	}

	b.addItem(item, priceFloat)
}

func updateTip(b *bill) {
	reader := bufio.NewReader(os.Stdin)

	tip := getInput("Enter tip: ", reader)

	// convert tip to float
	tipFloat, err := strconv.ParseFloat(tip, 64)

	if err != nil {
		fmt.Println("Invalid tip")
		return
	}

	b.updateTip(tipFloat)
}

func promptOptions(b *bill) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Options:")
	fmt.Println("\t1. Add item")
	fmt.Println("\t2. Update tip")
	fmt.Println("\t3. Print bill")
	fmt.Println("\t4. Save bill")
	fmt.Println("\t5. Exit")

	input := getInput("Enter your choice: ", reader)

	switch input {
	case "1":
		addItem(b)
		promptOptions(b)
	case "2":
		updateTip(b)
		promptOptions(b)
	case "3":
		fmt.Println(b.format())
		promptOptions(b)
	case "4":
		fmt.Println("Saving bill...")
		b.save(b.name)
		promptOptions(b)
	case "5":
		fmt.Println("Exiting...")
		os.Exit(0)
	default:
		fmt.Println("Invalid input")
		promptOptions(b)
	}
}

func main() {

	myBill := createBill()
	promptOptions(&myBill)

}
