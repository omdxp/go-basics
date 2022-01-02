package main

import "fmt"

func updateName(x string) string {
	x = "Sam"
	return x
}

func updateMenu(y map[string]float64) {
	y["Hamburger"] = 7.50

}

func main() {
	// group A types -> strings, ints, floats, bools, arrays, structs
	name := "Colby"

	name = updateName(name)

	println(name)

	// group B types -> slices, maps, functions
	menu := map[string]float64{
		"Hamburger": 4.50,
		"Hot Dog":   3.50,
		"Fries":     2.50,
	}

	updateMenu(menu)
	fmt.Println(menu)
}
