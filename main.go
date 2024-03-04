package main

import (
	"errors"
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

	runeShort := 'e'

	fmt.Println(runeShort)

	var myBoolean bool = true

	fmt.Println("Boolean: ", myBoolean)

	var someType = "jrklg"

	fmt.Println(someType)

	var1, var2, var3 := "jhjj", "jkd", 3

	fmt.Println(var1, var2, var3)

	const myConst string = "Something!!"

	print(myConst)

	print(fmt.Sprint(intDivision(2, 3)))

	result, remainder, error := intDivision(3, 8)

	if error != nil {
		print("Error encountered!")
	} else if remainder == 0 {
		print("Perfect division")
	}

	fmt.Printf("Dividing %v by %v, result is %v and remainder is %v", 8, 3, result, remainder)
}

func print(print string) {
	fmt.Println(print)
}

func intDivision(denom int, numer int) (int, int, error) {

	if denom == 0 {
		var error error = errors.New("division by zero")
		return 0, 0, error
	}
	return numer / denom, numer % denom, nil
}
