package main

import "fmt"

func main() {

	x := 0

	for x < 5 {
		fmt.Println(x)
		x++
	}

	for y := 0; y < 5; y++ {
		fmt.Println(y)
	}

	for {
		fmt.Println("loop")
		break
	}

	names := [4]string{"John", "Paul", "George", "Ringo"}

	for i, name := range names {
		fmt.Println(i, name)
	}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for _, name := range names {
		fmt.Println(name)
	}

}
