package main

import "fmt"

func goConditions() {
	//age := 50
	names := []string{"Ben", "Gerald", "Fiona", "Jane", "Kiera"}

	// if age < 30 {
	// 	fmt.Println("Age is less than 30.")
	// } else if age < 40 {
	// 	fmt.Println("Age is less than 40.")
	// } else {
	// 	fmt.Println("Age is not less than 45")
	// }

	for index, value := range names {
		if index == 1 {
			fmt.Println("Continuing at Position", index)
			continue //If condition is true, do not execute code below but continue from the start of the loop
		}

		if index > 2 {
			fmt.Printf("Breaking at position %v", index)
			break //Breaks the loop completely
		}

		fmt.Printf("The value at the position %v is %v \n", index, value)
	}
}
