package main

import "strings"

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	if len(names) == 1 {
		return names[0][0:1], "_"
	}
	return names[0][0:1], names[1][0:1]
}

func main() {

	i1, i2 := getInitials("tifa lockhart")
	println(i1, i2)

	i1, i2 = getInitials("tifa")
	println(i1, i2)

}
