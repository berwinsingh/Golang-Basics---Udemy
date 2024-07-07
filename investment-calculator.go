package main

import (
	"fmt"
	"math" //Since in GO there isn't a standard library that allows us to do the power of a number, we need to import the math library
)

// Creating an investment calculator and learning about other Go basics
func main() {
	const inflationRate float64 = 2.5 //The const keyword is used to declare a constant variable
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Enter the investment amount: ") //This print function does not add a new line at the end.
	fmt.Scan(&investmentAmount)                //The Scan function is used to read the input from the console thereby taking input from the user

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years) //Within the math.Pow function, the first argument is the base and the second argument is the exponent
	// var futureValue = float64(investmentAmount) * math.Pow(1+expectedReturnRate/100, float64(years)) //This is another way to write the same code as above however, we are converting the investmentAmount and years to float64 when being used
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println("The future value of the investment is:", futureValue) //The Println function is used to print the output to the console but in a new line
	fmt.Println("The future real value of the investment after inflation is:", futureRealValue)
}
