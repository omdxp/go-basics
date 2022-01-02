package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	greeting := "hello there friends"

	fmt.Println(strings.Contains(greeting, "hello!"))
	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "th"))
	fmt.Println(strings.Split(greeting, " "))

	// the original string is not modified
	fmt.Println(greeting)

	ages := []int{45, 23, 12, 34, 56, 78, 90}

	sort.Ints(ages)
	fmt.Println(ages)

	index := sort.SearchInts(ages, 101)
	fmt.Println(index)

	names := []string{"john", "jane", "joe", "jim", "jill"}
	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "jane"))

}
