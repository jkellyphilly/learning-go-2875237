package main

import (
	"fmt"
)

const aConstant int = 64

func main() {
	var aString string = "Yo, it's Joel."
	fmt.Println(aString)
	fmt.Printf("The type of aString var is %T\n", aString)

	var anInteger int = 7
	fmt.Println(anInteger)
	fmt.Printf("The type of anInteger var is %T\n", anInteger)

	var anotherInteger int
	fmt.Println(anotherInteger)

	var anotherString = "This is another string." // implicitly set type to string
	fmt.Println(anotherString)
	fmt.Printf("The type of anotherString is %T\n", anotherString)

	myString := "This string ain't got no var keyword in its initialization!"
	fmt.Println(myString)
	fmt.Printf("Type of myString: %T\n", myString)

	fmt.Printf("Value of aConstant is: %v\n", aConstant)
}

/*
OUTPUT:

Yo, it's Joel.
The type of aString var is string
7
The type of anInteger var is int
0
This is another string.
The type of anotherString is string
This string ain't got no var keyword in its initialization!
Type of myString: string
Value of aConstant is: 64
*/
