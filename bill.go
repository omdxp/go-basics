package main

import (
	"fmt"
	"os"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// format the bill
func (b *bill) format() string {
	fs := "\tBill breakdown:\n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25s $%.2f\n", k+":", v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("\t%-10s $%.2f\n", "Tip:", b.tip)
	total += b.tip

	// add total
	fs += fmt.Sprintf("\t%-10s $%.2f\n", "Total:", total)

	return fs
}

// update tip
func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add item to bill
func (b *bill) addItem(item string, price float64) {
	b.items[item] = price
}

// save bill to file
func (b *bill) save(filename string) {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+filename+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Saved bill to file:", filename+".txt")
}
