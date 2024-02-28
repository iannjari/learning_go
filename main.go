package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	fmt.Println("Something")

	var integer0 int16 = 389

	fmt.Println("Signed integer: ", integer0)

	var integer1 uint16 = 38945

	fmt.Println("Un-signed integer: ", integer1)

	var float0 float32 = 389.89065

	fmt.Println("Float: ", float0)

	// cast

	fmt.Println("Cast float to int. Result: ", int(float0))

	// division
	fmt.Println("int/int e.g 3/2 results in rounded-down int. Result: ", 3/2)

	var my_string1 = "Hello World!"
	var my_string2 string = "Hello \n World!"

	fmt.Println("String 1: ", my_string1)
	fmt.Println("String 2: ", my_string2)

	fmt.Println(`
	This is
	a 
	string block or whareva`)

	fmt.Println("String len() and RuneCount are different as len is len in bytes: ")
	fmt.Println(`len("θ") is :`, len("θ"))

	fmt.Println(`utf8.RuneCountInString("θ") is :`, utf8.RuneCountInString("θ"))

	var myRune rune = 'u'

	fmt.Println("Rune: ", myRune)
}
