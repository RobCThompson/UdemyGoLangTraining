package main

import (
	"fmt"
	"strconv"
	"strings"
)

func intSlicetoString(slice []int) string {
	chars := []string{}
	for _, n := range slice {
		chars = append(chars, strconv.Itoa(n))
		fmt.Println(chars)
	}
	return strings.Join(chars, "")
}

func main() {
	// Example byte message
	msg := []byte{0x12, 0x09, 0xa3, 0xfd}
	// iterate over each member of msg
	for _, n := range msg {
		// convert the byte into a binary string representation
		bits := fmt.Sprintf("%08b", n)
		// convert that string back to an int64
		i, _ := strconv.ParseInt(bits, 2, 16)
		// Print it.  Note conversino to byte.
		fmt.Printf("%#x\t%v\t%#x\n", n, bits, byte(i))
	}

	// Slice of ints
	newBits := []int{1, 0, 0, 1, 1, 1, 0, 1}
	fmt.Println()
	fmt.Println(newBits)
	newString := intSlicetoString(newBits)
	fmt.Println(newString)
	// 2 is the base, 16 is the output data width. Not sure why 8 doesn't work
	newInt, _ := strconv.ParseInt(newString, 2, 16)
	// newInt is int64, here convert (cast) int64 to byte
	newByte := byte(newInt)
	// Done. Print it.
	fmt.Printf("%T, %v\t%#x\n", newByte, newByte, newByte)
}
