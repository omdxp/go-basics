package main

import (
	"fmt"
	"math"
)

func sayGreeting(n string) {
	println("Greetings!", n)
}

func sayBye(n string) {
	println("Bye!", n)
}

func cycleNames(n []string, f func(string)) {
	for _, name := range n {
		f(name)
	}

}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func main() {

	sayGreeting("World")
	sayBye("World")

	cycleNames([]string{"World", "Gopher"}, sayGreeting)

	a := circleArea(5)
	println(a)
	fmt.Printf("%.3f\n", a)

}
