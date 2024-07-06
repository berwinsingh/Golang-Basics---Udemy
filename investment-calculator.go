package main

import (
	"fmt"
	"math" //Since in go there isn't a standard library that allows us to do the power of a number, we need to import the math library
)

func main() {
	// fmt.Println("Hello, World!") // Print Hello, World!

	//Creating an investment calculator and learning about other Go basics
	var investmentAmount, years float64 = 1000, 10
	var expectedReturnRate float64 = 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years) //Within the math.Pow function, the first argument is the base and the second argument is the exponent
	// var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years)) //This is another way to write the same code as above however, we are converting the investmentAmount and years to float64 when being used
	fmt.Println("The future value of the investment is:", futureValue) //The Println function is used to print the output to the console but in a new line
}
