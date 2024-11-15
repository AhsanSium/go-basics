// run command go run main.go
package main

import "fmt"

func main() {
	var nameOne string = "AHSAN"
	var nameTwo = "Luigi"
	var nameThree string

	fmt.Println(nameOne, "\n", nameTwo, "\n", nameThree)

	nameThree = "Bowser"

	fmt.Println(nameOne, "\n", nameTwo, "\n", nameThree)

	nameFour := "Mario"

	fmt.Println(nameOne, "\n", nameTwo, "\n", nameThree, "\n", nameFour)

	// Integers
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// Integer Bits
	var numberOne int = 10             // basic integer
	var numberTwo int8 = 127           // 8-bit integer
	var numberThree int16 = -32767     // 16-bit integer
	var numberFour uint32 = 2147483647 // 32-bit unsigned integer

	fmt.Print(numberOne, numberTwo, numberThree, numberFour)

	// Printing and Formatting
	fmt.Print("\nHello, ")
	fmt.Print("World\n")
	fmt.Println("Automatically Adding a New Line")
	fmt.Println("My age is", ageTwo, "And my name is", nameTwo)

	// Formatted String man
	fmt.Printf("First Number %v and Second Number %v \n", numberOne, numberTwo) // var in general
	fmt.Printf("First Name %q and Second Name %q \n", nameOne, nameTwo)         // for quoted string
	fmt.Printf("Type %T \n", numberFour)                                        // for type of variable
	fmt.Printf("Your Scored Point %0.2f \n", 22.22)                             // for type of variable

	// Save Printf
	var stringOne = fmt.Sprintf("Name is %s Age is %d \n", nameOne, ageOne)
	fmt.Println(stringOne)
}
