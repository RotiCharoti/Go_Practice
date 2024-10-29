//This file is a regular file, not an entry file
//Multiple files under package main will be considered to be
//part of the same Go program and executable

package main

import "fmt"

//Manual vars can be declared outside a function
//Can still be used and printed is a function below it
//This means that these vars are accessible on a package level, meaning other funcs 
//can also use it
var nameFive string = "Five"
var nameSix = "Six"

//Cannot be declared outside of function
//nameSeven := "Seven"
//Cannot use print statements outside of function
//fmt.Println(nameSixcan)

//For other functions to run, you must declare it in main function
//golang variables
//----------------------------------------------------------------------------------
func goStrings() {
	//strings
	var nameOne string = "Mario" //Error will occur when var is declared and not used
	var nameTwo = "Luigi"        //go infers from the value the data type of the var

	//go allows you to store data type of var without assigning a value to it as long as you print the empty value
	//Once var type are declared, it cannot be changed later on, unless you delete the first assigned data type
	var nameThree string
	fmt.Println(nameOne, nameTwo, nameThree)

	//Updating the values of vars
	nameOne = "Peach"
	nameThree = "Bowser"
	fmt.Println(nameOne, nameTwo, nameThree)

	//Initalize and declaring vars
	nameFour := "Yoshi"

	fmt.Println(nameFour)
	fmt.Println(nameFive, nameSix)
}

//----------------------------------------------------------------------------------
func goIntAndFloats() {
	//ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	//This line returns error as uint does not allow negative integers
	//var ageFour uint = -90

	fmt.Println(ageOne, ageTwo, ageThree)

	var scoreOne float32 = 90.989
	var scoreTwo float64 = 900.0039 //Not much diff between float32 and float64
	scoreThree := 89.9              //Default usage is float64
	fmt.Println(scoreOne, scoreTwo, scoreThree)
}

//----------------------------------------------------------------------------------
func goPrint() {
	age := 45
	name := "Jane"

	age2 := 19
	name2 := "Kenny"

	//Print
	fmt.Print("Hello ")   //Print does not add a new line at the end of the string
	fmt.Print("World,\n") //\n adds a new line to Print at the end of the string
	fmt.Print("Ninjas! ")

	//Println
	fmt.Println("My name is ", name)
	fmt.Println("And my age is ", age)

	//Printf (Formatted strings)
	//Format specifier %[data_type]
	//%v - vars
	//%q - quotes
	//%T - data type of the var
	//%f - float
	fmt.Printf("My age is %v and my name is %v \n", age, name)  //Does not add new line
	fmt.Printf("My name is %q and my age is %q\n", name2, age2) //Does not work on ints
	fmt.Printf("var name date type: %T | var age data type is: %T \n", name, age)
	fmt.Printf("You scored %f \n", 234.34)

	//Sprintf (Save formatted strings)
	var str = fmt.Sprintf("My age is %v and my name is %v", age, name) //Save the formatted string in a var using Sprintf
	var str2 = fmt.Sprintln("My age is 24 and my name is Mary")        //Save the string in a var
	fmt.Println("The saved string is: ", str)
	fmt.Println("The second saved string is: ", str2)
}

func goArrays() {
	var ages [3]int   //The var ages contains 3 items with data type int
	fmt.Println(ages) //Default value in an arrary with no assigned values is 0, hence output is [0,0,0]

	var values [3]int = [3]int{20, 23, 90} //no. of items an data type must be assigned to the right too when assigning values
	var values2 = [3]int{89, 43, 70}
	values[0] = 7777
	fmt.Println(values, values2)
	fmt.Println(values, len(ages))

	names := [2]string{"Peach", "Apple"}
	fmt.Println(names, len(names))
}

func goSlices() {
	//Slices use arrays under the hood
	var scores = []int{100, 50, 60, 21} //empty [] means no fixed no. of items in array, this is known as slices
	scores2 := []int{101}
	scores[1] = 10
	fmt.Println(scores, scores2)

	//Items can be appended to slices but not arrays
	scores = append(scores, 85) //85 is added to the slice
	fmt.Println(scores)

	//Slice ranges
	rangeOne := scores[1:3]  //1:3 means index 1 to index 2, as it stops before 3, inclusive of first index but not last
	rangeTwo := scores[2:]   //Inclusive of the last index
	rangeThree := scores[:3] //Not inclsuive of the last index
	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, 9000)
	fmt.Println(rangeOne)
}



//Good to know:
//var values can always be updated, but data types cannot be changed
//...ln usually prints a new line at the end of the string
//vars with same name but in diff function can still be executed and declared differently
//Slices can always be appended but not arrays
