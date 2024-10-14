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

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func (b *bill) format() string {
	fs := "Bill Breakdown: \n"
	var total float64 = 0

	//list items

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v)
		total += v
	}
	//tip
	fs += fmt.Sprintf("%-25v ...$%v \n", "tip:", b.tip)
	total += b.tip
	//total

	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)
	return fs
}

//update tip

func (b *bill) updateTip(tip float64) {
	b.tip = tip
}

// add item to the bill

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price
}

// save bill
func (b *bill) save() {
	// Check if the directory exists
	if _, err := os.Stat("bills"); os.IsNotExist(err) {
		// If the directory does not exist, create it
		err = os.Mkdir("bills", os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	// Turn the format into a byte slice
	data := []byte(b.format())
	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err) // panic stops the normal flow of the execution if there is any error and then print the err
	}
	fmt.Println("bill was saved to file")
}
