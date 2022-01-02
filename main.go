package main

import "fmt"

func main() {

	age := 18

	if age >= 18 {
		println("You are old enough to vote!")
		println("Have you registered to vote yet?")
	} else {
		println("Sorry, you are too young to vote.")
		println("Please register to vote as soon as you turn 18!")
	}

	number := 25

	if number < 30 {
		println("Number is less than 30")
	} else if number < 40 {
		println("Number is less than 40")
	} else {
		println("Number is greater than 40")
	}

	names := []string{"Bob", "Bill", "Joe", "Jim"}

	for i, name := range names {
		if i == 1 {
			println("continuing at pos", i)
			continue
		}

		if i > 2 {
			println("breaking at pos", i)
			break
		}

		fmt.Printf("the value at %d is %s\n", i, name)
	}

}
