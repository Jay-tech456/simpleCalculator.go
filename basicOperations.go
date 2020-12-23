package main
// Manjesh Prasad
// December 2, 2020

import (
	"fmt"							// used for Printing and scanning
)

func main(){
	num1 := 0
	num2 := 0
	for count:= 0; count < 2; count++{			// setting a for loop to allow the program
								// to run twice
	
		fmt.Printf("Please enter a value for num1: ")	// allow users to enter num1
		fmt.Scan(&num1)
		
		fmt.Printf("Pleae enter a value for num2: ")	// allow users to enter num2
		fmt.Scan(&num2)

			// addition
		fmt.Printf("%v + %v =  %v\n",num1, num2, num1 + num2)
		
			// subtraction
		fmt.Printf("%v - %v =  %v\n",num1, num2, num1 - num2)
		
			// multiply)
		fmt.Printf("%v * %v = %v\n", num1, num2, num1 * num2)
		
			// divide
		if (num2 != 0 ) {						// if b == 0, then we cant divide
			fmt.Printf("%v / %v = %v \n",num1, num2, num1 / num2)
		} else {
			fmt.Printf("Undefined")
		}
			// mod - will give you the remainder
		fmt.Printf("%v mod %v = %v\n",num1, num2, num1 % num2)

	}
									// end of program
}
		
