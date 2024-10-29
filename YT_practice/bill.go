package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type bill struct { //struct for bill object
	name  string
	items map[string]float64
	tip   float64
}

// Make new bill objects
func newBill(name string) bill { //Return bill object
	b := bill{ //Create bill object and initialize values
		name: name,
		items: map[string]float64{
			"soup":   4.99,
			"pie":    7.87,
			"salad":  5.77,
			"coffee": 8.99,
		},
		tip: 0,
	}

	return b
}

// Method
// Format the bill
func (b bill) format() string { //Receives a bill object and returns a string
	fs := "Bill breakdown: \n"
	var total float64 = 0

	//List items
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v) //-25 adds spacing to the right
		total += v
	}

	//Total
	fs += fmt.Sprintf("%25v ...$%0.2f", "Total: ", total) //(plus) 25 adds spacing to the left
	return fs
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	return strings.TrimSpace(input), err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput("Choose Option (a - add item, s -save bill, t - add tip): ", reader)

	//Control structures 
	switch opt {
	case "a":
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		// p, err := strconv.ParseFloat(price, 64)
		// if err != nil {
		// 	fmt.Println("The price must be a number.")
		// 	promptOptions(b)
		// }
		fmt.Println(name, price)
	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)
		fmt.Println(tip)
	case "s":
		fmt.Println("You chose s")
	default:
		fmt.Println("That was not a valid option...")
		promptOptions(b) //Repeats the entire function again
	}
}

func createBill() bill { //Creates a function that returns a bill object
	reader := bufio.NewReader(os.Stdin) //Creates a buffered reader that reads from standard input (os.Stdin), which is the terminal
	// fmt.Print("Create a new bill name: ")
	// name, _ := reader.ReadString('\n')
	// name = strings.TrimSpace(name) //Removes extra spaces in input

	name, _ := getInput("Create a new bill name: ", reader) //getInput returns 2 values, string and error, _ is used to not receive the error
	b := newBill(name)                                      //Creates bill object with input name
	fmt.Println("Created the bill -", b.name)

	return b
}

//go does not use classes but custom types using a struct
//Struct is a blueprint to describe a type of data
