/*
Quick demo of using the blank identifier to allow us to ignore returned variables.
In this case it's by avoiding the error checkig which actually shouln't be done in production code but there you go.

It also demonstrates how rediculously easy it is to perfom an HTTP get!
*/
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("http://www.google.co.uk")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s\n\n", page)
}
