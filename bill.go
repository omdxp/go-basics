package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make new bills
func newBill(name string) bill {
	b := bill{
		name: name,
		items: map[string]float64{
			"Food":   10.00,
			"Drinks": 5.00,
		},
		tip: 0,
	}
	return b
}

// format the bill
func (b bill) format() string {
	fs := "\tBill breakdown:\n"
	var total float64 = 0

	// list items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25s: $%.2f\n", k, v)
		total += v
	}

	// add tip
	fs += fmt.Sprintf("\tTip: $%.2f\n", b.tip)
	total += b.tip

	// add total
	fs += fmt.Sprintf("\tTotal: $%.2f\n", total)

	return fs
}
