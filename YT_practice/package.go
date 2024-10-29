package main

import (
	"fmt"
	"strings"
)

func goStringPackage() {
	greeting := "Hello there friends!"
	//strings.Contains returns true or false
	fmt.Println(strings.Contains(greeting, "Hello")) //Takes in 2 arguements, the string you want to use to search smth and the second is the search term

	//Returns the string with the newly replaced word but does not change the original value
	fmt.Println("This returns the string with replaced word:", strings.ReplaceAll(greeting, "Hello", "Hi")) //Takes in 3 arguements, the string, the word to replace, the word that is replacing the original word
	fmt.Println(strings.ToUpper(greeting))

	//Splits the element starting from "Hello", this returns [ there friends], removing the "Hello"
	fmt.Println(strings.Split(greeting, "Hello")) //Takes in 2 arguements, string, and the element to start the splitting from

	//Finds the index where the second arguement starts
	fmt.Println(strings.Index(greeting, "ll")) //Takes in 2 arguements, the string and the desired string to find
	fmt.Println("This is the original string, which remains unchanged:", greeting)

	//Sorts ints in ascending order
	// ages := []int{30, 10, 60, 89, 8, 90, 23}
	// sort.Ints(ages) //Cannot be stored in a var
	// fmt.Println(ages)

	//Returns the index of the int that is in second arguement
	//If index of second arguement does not exists, go returns the index after the last element in the slice
	// index := sort.SearchInts(ages, 30)
	// fmt.Println(index)
}
