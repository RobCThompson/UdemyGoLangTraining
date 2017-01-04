// Number printing exercise from lesson 35 of Go course at Udemy

// %b	base 2
// %c	the character represented by the corresponding Unicode code point
// %d	base 10
// %o	base 8
// %q	a single-quoted character literal safely escaped with Go syntax.
// %x	base 16, with lower-case letters for a-f
// %X	base 16, with upper-case letters for A-F
// %U	Unicode format: U+1234; same as "U+%04X"
package main

import "fmt"

func main() {
	// Print the number 42
	fmt.Println(42)
	// Print the number 42 in various formats
	fmt.Println("Decimal\tBinary\t\tOctal\tHexadecimal")
	fmt.Printf("%d\t%b\t%#o\t% x\n", 42, 42, 42, 42)

	// Now do it in a loop for the first 200 UTF-8 characters
	for i := 0; i < 200; i++ {
		fmt.Printf("%d  \t %08b \t %#o \t %#x \t %q\n", i, i, i, i, i)
	}
}
