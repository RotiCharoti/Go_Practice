package main

import "fmt"

func goLoop() {
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("The number is now ", i)
	// }

	names := []string{"Ben", "Gerald", "Fiona", "Jane", "Kiera"}

	//By using range, for each loop, go returns the index and the value of the index
	for index, value := range names {
		fmt.Printf("The value at index %v is %v \n", index, value)
	}

	//This prints out the index instead of value
	//By default, the first arguement returns index instead of value
	for value := range names {
		fmt.Printf("The value is %v \n", value)
	}

	//To print out value only
	for _, value := range names {
		fmt.Printf("The value is %v \n", value)
	}

	//To print out index only
	// for index, _ := range names {
	// 	fmt.Printf("The index is %v \n", index)
	// }
}
