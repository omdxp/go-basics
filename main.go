package main

func updateName(x *string) {
	*x = "Sam"
}

func main() {

	y := "Colby"
	updateName(&y)
	println(y)
	println("memory address of y:", &y)

}
