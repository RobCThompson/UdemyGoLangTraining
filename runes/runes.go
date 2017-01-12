package main

import "fmt"

func main() {
	myRune := '£'
	fehu := 0x16a0                      // Fehu, first rune of the elder fulthark
	fmt.Println("My rune: ", myRune)    // My rune: 163
	fmt.Printf("My rune: %q\n", myRune) // My rune: '£'
	fmt.Printf("Fehu: %q\n", fehu)      // My rune: '£'

	// Print the Futhark
	elderFuthark := []int32{0x16a0,
		0x16a2,
		0x16a6,
		0x16a8,
		0x16b1,
		0x16b2,
		0x16b7,
		0x16b9,
		0x16ba,
		0x16be,
		0x16c1,
		0x16c3,
		0x16c7,
		0x16c8,
		0x16c9,
		0x16ca,
		0x16cf,
		0x16d2,
		0x16d6,
		0x16d7,
		0x16da,
		0x16dc,
		0x16de,
		0x16df}
	for i := 0; i < len(elderFuthark); i++ {
		fmt.Printf("%v %U %q\n", elderFuthark[i], elderFuthark[i], elderFuthark[i])
	}
}
