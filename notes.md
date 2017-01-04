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
