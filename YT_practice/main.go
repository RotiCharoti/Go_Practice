package main

//import libraries

// You can assign values in the params
// func sayGreeting(n string) {
// 	fmt.Printf("Good Morning %v \n", n)
// }

// func sayBye(n string) {
// 	fmt.Printf("Goodbye %v \n", n)
// }

// func cycleName(n []string, f func(string)) { //Pass a string slice and a function that takes a string as the param
// 	for _, g := range n {
// 		f(g)
// 	}
// }

// func circleArea(r float64) float64 { //to return a value, you have to specify the data type of the value
// 	return math.Pi * r * r
// }

// Does not matter if function main is placed at the top or not,
// as go reads the entire program at once instead of line by line
func main() {
	//goStrings()
	//goIntAndFloats()
	//goPrint()
	//goArrays()
	//goSlices()
	//goStringPackage()
	//goLoop()
	//goConditions()
	//goMap()
	mybill := createBill() //Has a print output and returns bill object
	promptOptions(mybill)

	//mybill := newBill("Jane's bill")
	//fmt.Println(mybill.format())

	// fn1, sn1 := getIntials("cloud nine")
	// fmt.Println(fn1, sn1)
	// fn2, sn2 := getIntials("bae")
	// fmt.Println(fn2, sn2)

	// sayGreeting("Jane")
	// sayGreeting("Mary")
	// sayBye("Vary")
	// names := []string{"cloud", "tifa", "barret"}
	// cycleName(names, sayGreeting)
	// cycleName([]string{"cloud", "tifa", "barret"}, sayBye)

	// a1 := circleArea(10.5)
	// a2 := circleArea(15.982)
	// fmt.Println(a1, a2)
	// fmt.Printf("Circle 1: %0.3f\nCircle 2: %0.3f", a1, a2) //To round off to decimal places, use %0.[no. of places to round off to]f
}

//There must be only one main function in the entry file, located within the main package
//Other functions can be created in other files and invoked in the entry file
//main function executes all the other functions in the package
//You can call the function here and pass values in the params

//Good to know:
//Common practice to call entry file main.go
//Type in go run . to run your functions from another file that is in main file
//or go run main.go [name_of_your_other_file].go
//Can only be done after typing in go mod init [name_of_directory]
