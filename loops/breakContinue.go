package main

import "fmt"

func main() {
	i := 0
	for {
		i++
		if i%2 == 0 {
			continue // i % 2 == 0 for EVEN numbers, in that case contiune to the next
			// loop iteration.
		}
		fmt.Println(i)
		if i >= 50 {
			break // Breaks on 51 because 50 is even.  In that case the continues
			// prevents this block from being reached.
		}
	}
}
