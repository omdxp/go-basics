package main

import "fmt"

func main() {

	menu := map[string]float64{
		"Hamburger":    2.5,
		"Cheeseburger": 3.5,
		"Hotdog":       2.0,
		"Fries":        1.5,
		"Drink":        1.0,
	}

	fmt.Println(menu)
	fmt.Println(menu["Hamburger"])

	// looping maps
	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	// ints as key type
	phonebook := map[int]string{
		1: "John",
		2: "Jane",
		3: "Jack",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[1])

	phonebook[2] = "Jill"

	fmt.Println(phonebook)

}
