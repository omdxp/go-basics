package main

import "fmt"

func main() {

	// strings
	var nameOne string = "Omar"
	var nameTwo = "Yasser"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "Peach"
	nameThree = "Kiwi"

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Pineapple"
	fmt.Println(nameFour)

	// integers
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory
	var numOne int8 = 25
	var numTwo int8 = -128
	var numThree uint8 = 255

	var scoreOne float32 = 25.98
	var scoreTwo float64 = 2332352267989608.75445
	scoreThree := 1.5

}
