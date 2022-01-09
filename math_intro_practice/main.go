package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 12, 45, 48
	intSum := i1 + i2 + i3

	fmt.Println("Integer sum is: ", intSum)

	f1, f2, f3 := 2.0 + 3.5 + 8.229
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum: ", floatSum)

	// safest way of holding on to precision
	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("The sum is now:", floatSum)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi         // 2*pi*R
	fmt.Printf("Circumference: %0.2f\n", circumference) // 0.2f asks for 2 digits after decimal point
}
