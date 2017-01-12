package main

import "fmt"

func main() {
	x := 13 / 4
	r := 13 % 4

	fmt.Printf("13/4 = %d  13%%4 = %d\n", x, r)

	y := 13.0 / 4
	fmt.Printf("13.0/4 = %v\n", y)
}
