package main

import (
	"fmt"
)

func goMap() {
	menu := map[string]float64{ //Map requires a [key] and a value
		"soup":   4.99,
		"pie":    7.87,
		"salad":  5.77,
		"coffee": 8.99, //Must end wiht a comma
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"]) //Prints out value of key

	//Looping maps
	for k, v := range menu { //k = key, v = value
		fmt.Println(k, "-", v)
	}

	//Ints as key type
	phonebook := map[int]string{
		89087574: "Jane",
		90276253: "Mary",
		90282335: "John",
		93056319: "Nathan",
	}

	fmt.Println(phonebook)
	fmt.Println(phonebook[89087574])

	phonebook[93056319] = "Bethany"
	fmt.Println(phonebook[93056319])
}
