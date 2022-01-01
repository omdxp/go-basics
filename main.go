package main

import "fmt"

func main() {

	// var ages [3]int = [3]int{10, 20, 30}
	var ages = [3]int{10, 20, 30}

	names := [3]string{"Omar", "Yasser", "Mario"}
	names[1] = "Luigi"

	fmt.Println(ages, len(ages), cap(ages))
	fmt.Println(names, len(names), cap(names))

	// slices
	var slice1 []int = ages[:]
	var slice2 []int = ages[1:]
	var slice3 []int = ages[:2]
	var slice4 []int = ages[1:2]

	fmt.Println(slice1, len(slice1), cap(slice1))
	fmt.Println(slice2, len(slice2), cap(slice2))
	fmt.Println(slice3, len(slice3), cap(slice3))
	fmt.Println(slice4, len(slice4), cap(slice4))

	slice4 = append(slice4, 85)
	fmt.Println(slice4, len(slice4), cap(slice4))

}
