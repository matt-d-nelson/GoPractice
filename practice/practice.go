package practice

import (
	"fmt"
)

func HelloWorld() {
	fmt.Println("hello world")
}

/*
Complete the solution so that it splits the string into pairs of two characters.
If the string contains an odd number of characters then it should replace the
missing second character of the final pair with an underscore ('_').
*/

func Solution(str string) []string {
	var res []string
	if len(str)%2 != 0 { //check to ensure the string is even
		str += "_" // if not, add a character
	}
	for i := 0; i < len(str); i += 2 {
		res = append(res, str[i:i+2])
	}
	return res
}
