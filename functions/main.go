package main

import "fmt"

func greet(fname, lname string) string {
	return fmt.Sprint("Hello", fname, lname)
}

func main() {
	fmt.Println(greet("Rob", "T"))
	n := average(43, 56, 87, 12, 45, 57)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("sf: %v of type %T\n", sf, sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
