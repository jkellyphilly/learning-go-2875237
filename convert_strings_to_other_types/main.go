package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a name: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)

	fmt.Print("Enter a number: ")
	numInput, _ := reader.ReadString('\n')

	aFloat, err := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value of number:", aFloat)
	}
}

/*
CONSOLE OUTPUT (with my input):

Enter a name: Joel
You entered: Joel

Enter a number: 42.1902
Value of number: 42.1902

AGAIN, if entering not a number:
Enter a number: 8fja
strconv.ParseFloat: parsing "8fja": invalid syntax
*/
