package main

import (
	"strings"
)

func getIntials(n string) (string, string) { //Returns 2 values that are both string type
	s := strings.ToUpper(n)        //Converts the input string to all uppercase
	names := strings.Split(s, " ") //Splits each element and stores into a string slice
	var initials []string
	for _, v := range names {
		initials = append(initials, v[:1]) //v[:1] returns the first letter of the name
	}

	if len(initials) > 1 {
		return initials[0], initials[1] //Return values
	}

	return initials[0], "_"
}

//Further explanation of v[:1], which is [0:1]
//If John is in initials slice, v[:1] includes all the index but 1
//hence J[0] o[1] h[2] n[3] returns J
//The reason why [0] is not accepted is because go
//recognizes [0] as a byte, not a character
