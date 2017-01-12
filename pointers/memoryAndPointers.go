package main

import "fmt"

func main() {

	a := 43 // A normal variable, called a with the value 43

	var b *int = &a // A pointer to an int varaible, assigneed to the memory address of a

	fmt.Printf("a = %d &a = %v b = %v *b = %v\n", a, &a, b, *b)
}
