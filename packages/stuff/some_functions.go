package stuff

import "fmt"

// MyName is a variable holding my name!
var MyName = "Rob"

// DoSomething prints the message "Ok, ok, I'm doing it!" to the console
func DoSomething() {
	fmt.Println("Ok, ok, I'm doing it!")
}

// WhoAreYou print MyName to the console
func WhoAreYou() {
	fmt.Println(MyName)
}
