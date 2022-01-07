package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter a name: ")
	input, _ := reader.ReadString('\n')

	fmt.Println("You entered:", input)
}

/*
CONSOLE OUTPUT (with my input):

Enter a name: Joel
You entered: Joel
*/
