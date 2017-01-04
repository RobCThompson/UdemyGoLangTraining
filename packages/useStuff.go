package main

import (
	"fmt"

	"github.com/RobCThompson/UdemyGoLangTraining/stuff"
)

func main() {
	// Print a helpful message
	fmt.Println("Ok, off we Go...")
	stuff.DoSomething()
	fmt.Printf("Here's my name: %v", stuff.MyName)
	// Who are you?
	stuff.WhoAreYou()
}
