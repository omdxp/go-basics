package main

import "fmt"

func main() {

	myBill := newBill("Omar's Bill")

	myBill.addItem("Coffee", 2.50)
	myBill.addItem("Cake", 5.00)
	myBill.addItem("Chips", 1.00)

	myBill.updateTip(.20)

	fmt.Println(myBill.format())

}
