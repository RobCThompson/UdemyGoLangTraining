package main

import "fmt"

func main() {
	age := 44
	fmt.Println(age)  //  Value of age: 44
	fmt.Println(&age) //  Address of age: 0xc8200761c0

	changeMe(&age) // call changeMe passing in the address of age

	fmt.Println(age)  // New value of age: 24
	fmt.Println(&age) // Address of age: 0xc8200761c0
}

func changeMe(z *int) {
	fmt.Println(z)  // Reference, address: 0xc8200761c0
	fmt.Println(*z) // Dereference, value at that address:
	*z = 24         // Change hte value at that address
	fmt.Println(z)  // Address: 0xc8200761c0
	fmt.Println(*z) // Value at that address: 24
}
