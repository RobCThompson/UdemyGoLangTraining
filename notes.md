# Notes from the Udemy Course
# Learn How to Code: Google's Go (golang)
Todd McLeod

## Useful Resources
https://goo.gl/PHKgO7

--------------------------------------------------------------------------------
# 02: Why Go?

Invented by Google.

Notable contributors:

* Ken Thompson (B, C, Unix, UTF-8)
* Rob Pike (Unix, UTF-8)
* Robert Griesemer (Hotspot, JVM)

> GO was invented by geniuses.

* High performance
* Multiple cores: Language written to leveage multiple cores easily.
* Concurrency: Not necessarily easier to code but "baked" into the runtime
* Compiled
  * can comile for other arcitechtures on one machine.
* Network

Basically, it does everything and loads of companies are using it.

--------------------------------------------------------------------------------
# 03: Hello World

~~~ golang
package main
import "fmt"

func main(){
  fmt.Println("Hello, world!")
}
~~~

We start with a package declaration which identifies which package this code
belongs to.

Import is used to introduce libraries of code into our own projects.

golang.org vs godoc.org

godoc has the standard library *and* 3rd party libraries, golang.org only covers
the standard library.

--------------------------------------------------------------------------------
# 04: Section Overview

--------------------------------------------------------------------------------
# 05: Terminal Emulator

if using Windows, download github for windows... other wise skip the video.

--------------------------------------------------------------------------------
# 06: Installation Insights

https://golang.org

## SHA1 checksum
Is a hash algorithm used to check that the download is unmodified.

~~~ bash
openssl sha1 path_to_downloaded_item
~~~

Verify the Go installation with

~~~
go version
~~~

--------------------------------------------------------------------------------
# 07: Go workspace

The Go workspace:

* Some folder
  * bin
  * pkg
  * src
    * github.com
      * <github.com username>
        * folder with code for project/repo
        * folder with code for project/repo
        * folder with code for project/repo

--------------------------------------------------------------------------------
# 08: Go PATH Variables

--------------------------------------------------------------------------------
# 09: How to configure Environment Variables on Windows

On Windows7 Environment Variables are under "Advanced System Settings"

--------------------------------------------------------------------------------
# 10: How to configure Environement Variables on Mac

In terminal, head to user's root and edit `.bash_profile` (or `.bashrc`)

* $GOROOT should be pointed to `/usr/local/go`
* $GOPATH should be where ever you do your work, e.g `~/Documents/Coding/Go`
* add to $PATH the bin directory so your binaries are available everywhere.
  * `export PATH="$HOME/Documents/Coding/Go/bin:$PATH"`

--------------------------------------------------------------------------------
# 11 to 14: Configuring Linux installatino

  Skipped for now.

--------------------------------------------------------------------------------
# 15: Testing your installation
 Use `go get` to download some repository for use later.

--------------------------------------------------------------------------------
# 17: Development Environment - Preview


--------------------------------------------------------------------------------
# 18 & 19: GO Editors

There's a presentation on this at < https://goo.glPHKgO7 >

This guy *loves* WebStorm...

For Atom we will need the **go-plus** plugins

--------------------------------------------------------------------------------
# 22: Hello World... again?

~~~ golang
package main

import "fmt"

func main(){
  fmt.Println("Hello. world!")
}
~~~

--------------------------------------------------------------------------------
# 23: go Commands

`go help` will yield a list of go commands.

* `go run` is used to run
* `go build <filename>` builds the project and puts the executable into the working src folder.  The executatble will have the same name as the source file.
* `go build` build all the project file in the current folder. THe executable will have the name of the **containing** directory. So if you are in a directory `~/$GOPATH/src/github.com/RobCThompson/UdemyGoLangTraining` then the executable will be called UdemyGoLangTraining.
* `go clean` removes executable files from a working directory. .go source files are left intact.
* `go install` works the same as go buld except that it places the executable into the $GOPATH/bin directory.

--------------------------------------------------------------------------------
# 24 & 25: GitHub

Make sure that Git is installed.

Make sure you are in the project directory and initialise Git with `git init`.

Add the remote Github origin (if you are using it) with `git remote add origin https://github.com/<username>/<repo>` e.g `git remote add origin https://github.com/RobCThompson/UdemyGoLangTraining.git`.  Don't forget the `.git` extension!.

If the repository was created on Github with a README.md document or some other content then the repo's might not match so it's likely to be a good idea to *PULL* the remote repo before making changes to the local one.  Do this with `git pull origin master`.

Then, continue to use git as you would locally, i.e adding files with `git add <filename>` or `git add -all` and commiting with `git commit -m "a useful message"`.

If you then wish to push back to github do so with `git push -u origin master` or just `git push`.

Don't forget that `git status` is a useful command!

--------------------------------------------------------------------------------
# 27, 28, 29: Computing Fundamentals

--------------------------------------------------------------------------------
# 30: Housekeeping

Tutorial repo has changed to `github.com/goestoeleven/golangtraining`.

Can get this with Go using `go get -u github.com/goestoeleven/golangtraining`

--------------------------------------------------------------------------------
# 32: Binary stuff

Base 2

| 128's | 64's  | 32's  | 16's  | 8's   | 4's   | 2's   | 1's   |
|-------|-------|-------|-------|-------|-------|-------|-------|
| 2^7   |  2^6  | 2^5   | 2^4   | 2^3   |  2^2  |  2^1  | 2^0   |

--------------------------------------------------------------------------------
# 33: Hexadecimal Stuff
Base 16

|          | 65,536's |  4,096's | 256's    |  16's    |    1's   |
|----------|----------|----------|----------|----------|----------|
|          |   16^4   |   16^3   |    16^2  |    16^1  |   16^0   |
| 42 =     |          |          |          |     2    |    a     |

Why hexadecimal? More efficient storage.

## There's also Octal...
Base 8

| 262,144's | 32,768's  | 4096's  | 512's | 8's   |  64's | 8's   | 1's   |
|-------|-------|-------|-------|-------|-------|-------|-------|
| 8^7   |  8^6  | 8^5   | 8^4   | 8^3   |  8^2  |  8^1  | 8^0   |

--------------------------------------------------------------------------------
# 35: Coding Scheme Programs

Read <https://godoc.org/fmt> for format specifiers.
Print the number 42.

~~~golang
package main

import "fmt"

func main() {
	// Print the number 42
	fmt.Println(42)
	// Print the number 42 in various formats
	fmt.Println("Decimal\tBinary\t\tOctal\tHexadecimal")
	fmt.Printf("%d\t0b%08b\t%o\t0x%x\n", 42, 42, 42, 42)
}

~~~

Outputs:
~~~
42
Decimal    Binary        Octal    Hexadecimal
42         0b00101010    52       0x2a
~~~

--------------------------------------------------------------------------------
# 39: Packages

Go code is organised into packages. Folder names need to match package names.

For 3rd pary packages the import in the code needs to be the fully qulified name.

Functions with a Capital initial letter are exported, that is, visible from outside the package.  Functinos with a lower-case letter are only available inside the package.

Variables work in the same way!

So, in the project directory create a directory called "stuff" ad create a go source file in there.

~~~
mkdir stuff
touch stuff/some_functions.go
~~~

In the some_functions file:

~~~ golang
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
~~~

Now we can use these functions and variables.  Now, we can return to our project folder, create a program, say `touch useStuff.go` and use those functions and variables:

~~~ golang
package main

import (
	"fmt"

	"github.com/RobCThompson/UdemyGoLangTraining/stuff"
)

func main() {
	// Print a helpful message
	fmt.Println("Ok, off we Go...")
	stuff.DoSomething()
	fmt.Printf("Here's my name: %v\n", stuff.MyName)
	// Who are you?
	stuff.WhoAreYou()
}
~~~

--------------------------------------------------------------------------------
# 41: Variables

There are two preferred (idiomatic) ways to initialise variables:

* Shorthand: Can only be used inside a `func`.  Use `name := value`
* var: Can be used to initialise to the "zero" value

Declare, assign, initialise

Declare is where you create a variable with it's type but don't assicgn a value.

When printing we ca use %T to print the type of a variable.

--------------------------------------------------------------------------------
# 42: Scope

From where can I access that variable?

Rule of thumb - "keep your scope tight" - i.e. declare your variables close to where you are using them... and remember to create variables *before* trying to use them!

* **Package level**: When a variable is declared outside of a `func`.
* **Block Level**: When a variable is declared within a block say, a `func` or a `loop`, for example.

--------------------------------------------------------------------------------
# 46: Blank Identifier

You must use everything you put into your code.  All declared variables must be used.  The blank identifier is used to take any return values from functions that you don't need to use and divert them to nil.

Here's a program where we've ignored the error checking by dumping the err return values into the blank identifier.  It's also a pretty nifty example of how esy HTTP requests are in Go!

~~~ golang
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

~~~

--------------------------------------------------------------------------------
# 50, 51, 52: Memory Addresses and Pointers

Memory addresses can be accessed using the `&` operator.  Pointer's are identified by the `*` operator.  The same as in C/C++.

Memory addresses can be *referenced* by pointers:

~~~golang
a := 43         // a is an integer variable
var b *int = &a // b is a reference to the memory address of a
~~~

The values stored in the memory locations referenced by pointers can be accessed or *dereferenced* using the `*` operator.

~~~ golang
a:= 43
var b *int = &a
fmt.Printf("a=%v *b=%v", a, *b) // Prints: a=43 *b=43
~~~

## What's the point of pointers?
You can write a function to change a variable directly rather than using a return value.

--------------------------------------------------------------------------------
# 53: Remainder

Modulo.  Go's modulo operator is `%`.
Remember that division is integer so long as all the variables are integers.

~~~golang
x := 13 / 4   // 3
y := 13.0 / 4 // 3.25
~~~

--------------------------------------------------------------------------------
# 54: Review

See above!

Plus, you can generate/extract documentation from the source code using `godoc <package name>`.

For example:

`godoc ./packages/stuff`

--------------------------------------------------------------------------------
# 55: Section Overview

Check out <https://forum.golangbridge.org/> for asking questions.

Interesting note, instruction videos should be between 5 and 8 minutes long to provide ballance between viewers feeling they're "getting their money's worth" and fitting instide their attentino span.

--------------------------------------------------------------------------------
# 56: Loops

## Syntax: For Loop

~~~golang
for i:= 0; i <= 100; i++ {
  fmt.Println(i)
}
~~~

Syntax is more or less that same as C/C++.

--------------------------------------------------------------------------------
# 57: Nested Loops

 Wheen nesting loops remember that the "inner" loops loop through each time the outer loops loop through.

 ~~~golang

 for i := 0; i < 3; i++ {
   for j := 0; j < 3; j++ {
     for k := 0; k < 3; k++ {
       fmt.Println(i, " ", j, " ", k)
     }
   }
 }
~~~

Outputs:

~~~
0   0   0
0   0   1
0   0   2
0   1   0
0   1   1
0   1   2
0   2   0
0   2   1
0   2   2
1   0   0
1   0   1
1   0   2
1   1   0
1   1   1
1   1   2
1   2   0
1   2   1
1   2   2
2   0   0
2   0   1
2   0   2
2   1   0
2   1   1
2   1   2
2   2   0
2   2   1
2   2   2
~~~

--------------------------------------------------------------------------------
# 58: Loops Conditions, Break and Continue

In `for` loop the condition will cause the loop to evaluate while the condition evaluates to TRUE.

~~~golang
// Prints 0 to 9 to conole
i := 10
for i < 10 {
  fmt.Println(i)
  i++
}
~~~

~~~golang
// for with no condition - will run forever
i := 0
for {
  fmt.Println(i)
  i++
}
~~~

The following snippet prints 0 to 10, it uses a break statement in a condition to break out of the loop when the counter reaches 10.
~~~golang
// Prints 0 to 10.  Uses a
i := 0
for {
  fmt.Println(i)
  if i >= 10{
    break
  }
  i++
}
~~~

The keyword `continue` causes a loop to return to the top of the loop.  The following snippet only prints ODD numbers 1 to 51:

~~~golang
i := 0
for {
  i++
  if i %2 == 0 {
    continue  // i % 2 == 0 for EVEN numbers, in that case contiune to the next
              // loop iteration.
  }
  fmt.Println(i)
  if i>= 50{
    break     // Breaks on 51 because 50 is even.  In that case the continues
              // prevents this block from being reached.
  }
}
~~~

--------------------------------------------------------------------------------
# 59: Documentation and Terminology

In UTF-8 readable characters are encoded as numbers.

## Runes
**Rune**: A Rune is a single character from any language.

Strings are initialised using "" or ``, runes are initialised using single-quotes ''.
 Runes are also an alias for an `int32` since that's all a character is underneath.

 ~~~golang
 func main() {
 	myRune := '£'

 	fmt.Println("My rune: ", myRune)    // My rune: 163
 	fmt.Printf("My rune: %q\n", myRune) // My rune: '£'
 }
 ~~~

--------------------------------------------------------------------------------
 # 60: More on Runes

 In go *Casting* is called **Conversion**.

 For example, here we convert a string "Hello" to a slice of bytes:

 ~~~golang
 func main(){
   fmt.Println([]byte("Hello"))
 }
 ~~~

--------------------------------------------------------------------------------
# 61: Strings

 A string is just a collection of runes.

 Strings can be converted to slices of bytes using `[]byte("Hello")`.

 Strings can be declared inside double-quotes, "" or backticks ``.  Backticks allows the inclusion of quotes, new lines and other special characters.  Useful for declaring large blocks of text such as XML, HTML or SQL.

--------------------------------------------------------------------------------
# 62: Switch Statements

Switch statements in go don't automatically "fall through" like they do in C/C++ which means there's no need for the `break;` statement between cases.  To deal with scenarios when it's desirable to fall through then there is the `fallthrough` keyword.  A default, catch-all can be defined with the `default` keyword.

Multiple matches can have the same code by using commas `,` to separate expressions.


~~~golang
func main(){
  myScore := 3
  switch myScore {
  case 0, 1, 2:
    fmt.Println("Lame!")
  case 3,4,5:
    fmt.Print("Better luck next time!")
  case 6,7,8:
    fmt.Print("Hey, nice!")
  case 9:
    fmt.Println("Ooh, so close - good stuff!")
  case 10:
    fmt.Println("BOOM! Nailed it!")
  default:
    fmt.Println("nah, you're making it up!")
  }
}
~~~

The conditional statemenet can be omitted in favour of separate conditionals in the case statements to create if-else-if chains:

~~~golang
func main(){
  myScore := 3
  switch {
  case myScore <= 3:
    fmt.Println("Lame!")
  case myScore > 3 && myScore <= 5:
    fmt.Print("Better luck next time!")
  case myScore > 5 && myScore <= 8:
    fmt.Print("Hey, nice!")
  case myScore == 9:
    fmt.Println("Ooh, so close - good stuff!")
  case myScore == 10:
    fmt.Println("BOOM! Nailed it!")
  default:
    fmt.Println("nah, you're making it up!")
  }
}
~~~

--------------------------------------------------------------------------------
# 63: If

If Syntax in Go is pretty much the same as other languages.  THe block of colde runs if the statement evaluates to true.

The `!` is used to invert or negate the statement, read 'not', e.g `if a != b` "if a is not equal to b".

It's valid to include a variable initialisation in the if statement:

~~~golang
if err := file.Chmod(0664); err != nil{
  log.Print(err)
  return err
}
~~~

This is done to limit the scope of a variable.  It's scope is only the if block.

If-else syntax is:

~~~ golang

if false {
  fmt.Println("This will never print")
} else if false {
  fmt.Println("This won't print either")
} else if true {
  fmt.Println("This will print though")
} else {
  fmt.Println("Umm, else what? This won't print.")
}
~~~

--------------------------------------------------------------------------------
# 65: Section Review

--------------------------------------------------------------------------------
# 66, 67, 68: Functions

`func main(){}` is the entry point to a Go program, you can only have one!

In Go, parameters are defined with their name followed by their type (C/C++ et al do it the other way around).

~~~golang
func greet(name string){
  fmt.Println("Hello ", name)
}
~~~

## Returns

~~~golang
func greet(fname, lname string) string {
  return fmt.Sprint(fname, lname)
}

func main(){
  fmt.Println(greet("Rob", "Thompson"))
}
~~~

If you name the variable to return in the function declaration then the return doesn't need to be followed explicitly by the value to return:

~~~golang named returns
func greet(fname string, lname string) (s string)  {
  s = fmt.Sprint(fname, lname)
  return
}
~~~

Multiple return values can be defined in either way:

~~~golang
func greet(fname, lname string) (string, string){
  return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
~~~

## Variadic functions

Variadic functions can accept a variying number or arguments.

The final parameter in a functino signature my have a type prefixed with `...`.  A function with such a parameter is called *variadic* and may be invoked with zero or more arguments for that parameter.

~~~golang

func average(sf ...float64) float64 {
  fmt.Println(sf)
  fmt.Printf("sf: %v of type %T", sf, sf)
  var total float64
  for _, v := range sf {
    total += v
  }
  return total / float64(len(sf))
}

~~~

Using the `...` with an argument opens up the suppled slice and suplies the entities independently.

~~~golang
func main(){
  data := []float64{43, 56, 87, 12, 45, 57}
  n := average(data...)
  fmt.Println(n)
}
~~~

So, when declaring a function signature the dots for *before the paramater*.  When supplying data into a function as an argument the dots go *after the argument*.

## Func Expressions / Anonymous Functions

You can assign a function to a variable and then call it.

~~~golang

func main(){
  greeting := func() string {
    fmt.Println("Hello world!")
  }
  greeting()
  fmt.Printf("%T\n", greeting)
}
~~~

--------------------------------------------------------------------------------
# 72: Closure

All to do with scope.  Returning a function from a function is an example of closure...?

--------------------------------------------------------------------------------
# 73 & 74: Callbacks

In Go "function" is a type, therefor they can be passed around in the same wasy as any other variable.  They can be passed into functions and returned from functions.

This is an example of a callback, here a function, filter, accepts a slice of numbers and a function that takes an int as it's parameter and returns a boolean.  The filter function will return a slice containing numbers that have been processed by the function supplied to it.

In main() we create a variable that will hold the output of the filter function.  The slice and anonymous function are initialised at the same point as calling filter.

This *could* have been achieved without using a callback but in that instance the function's scope would have been package level.  It's a good thing to keep scope as tight as possible.
Also, callbacks are useful in event driven programming with languages such as Javascript, C# and C++ - is this also true for Go...?

~~~golang
package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
  var xs []int
  for _, n := range numbers {
    if callback(n){
      xs = append(xs, n)
    }
  }
  return xs
}

func main(){
  xs := filter([]int{1,2,3,4,5,6,7,8}, func(n int) bool {
    return n > 1
  })
  fmt.Println(xs) // [2 3 4 5 6 7 8]
}
~~~

--------------------------------------------------------------------------------
# 75: Recursion

When a function call itself then it is recursive.

Any problem that can be solved with recursino can be solved witout it.  Most of the time it's better just to use loops as they are more performant.

--------------------------------------------------------------------------------
# 76: Defer

Defer causes the folowing statement to be 'deferred' until the containing function is ready to exit.

In this example the defer statement causes the `world()` function to wait until just before the `main()` function exits which mean it happens **after** the `hello()` functino even though it's called first.

~~~golang
package main

import "fmt"

func hello() {
  fmt.Print("hello ")
}

func world() {
  fmt.Println("world")
}

func main() {
  defer world()
  hello
}

// Prints hello world
~~~

A good example is when working with files - it's important to remember to close files once we're done with them.  `defer` is useful because you can call the file close operation immediately after opening it but by deferring it the file only actually gets closed once you're other processing is done.
