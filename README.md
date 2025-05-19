# gocolor 🎨

A lightweight and didactic Go library for applying ANSI colors to your terminal output.
Originally designed as an academic and fun experiment, gcolor makes it easy to print colored messages in CLI applications using simple and expressive functions.

Whether you’re building a shell tool, debugging with styled logs, or just learning how ANSI escape codes work in Go, this library helps you get colorful fast — no dependencies, no complexity.

### ✨ Features

    • Minimal and easy-to-use API
    • Supports standard ANSI foreground colors (red, green, blue, yellow, etc.)
    • Built-in helpers like PrintRed("error"), CyanCor("info"), etc.
    • Struct-based coloring for flexible output formats
    • Safe and composable string building using gconcat
    • Useful for scripting, logging, and educational purposes


## 🧪 Simple Examples

```go
package main

import . "github.com/jeffotoni/gcolor"

func main() {
    PrintRed("This is an error message")
    PrintGreen("Operation successful!")
    Yellow.Cprintln("Running task...")
}
```

## Install

```go

go get github.com/jeffotoni/gcolor

```

## Use gcolor example

```go

package main

import (
	  "fmt"
	. "github.com/jeffotoni/gcolor"
)


//
//
//
type Jovem struct {
	Nome string
}

//
//
//
func (j Jovem) Fala() string {
	return fmt.Sprintf("Meu nome é : %s", j.Nome)
}

//
// start
//
func main() {

	//
	// First instantiate form
	//
	var C Collor
	C.Cor = "purple"
	C.Cprintln("Purple")

	//
	// Second way to instantiate
	//
	c := new(Collor)
	c.Cor = "green"
	c.Cprintln("Collor Green")

	//
	// Third way to instantiate
	//
	red := Collor{"red"}
	red.Cprintln("Collor red")

	//
	// Fourth way to instantiate
	//
	Red.Cprintln("Red color instance otherwise")

	//
	// Yellow
	//
	Yellow.Cprintln("Yellow color!!")

	//
	// Black Color
	//
	println(BlackCor("Testing black color"))

	//
	// Cyan Color
	//
	println(CyanCor("Testing cyan color"))

	//
	//
	//
	Texts := []string{"mytest 1", "mytest 2", "mytest 3"}

	//
	//
	//
	Printer.Print(YELLOW_FG, "%s", Texts)

	//
	//
	//
	j := Jovem{Nome: "Marcus Mann"}

	//
	//
	//
	cc := Color{}

	cyan := cc.Red(j.Fala())

	//
	//
	//
	fmt.Println(cyan)
}


```
