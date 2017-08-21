# gocolor

A simple lib to help you color when developing some bash application using ANSI-COLORS.

The goal is to make the lib leaner and leaner.

The gcolor was made only with this intúito of leaving its characters in the colored bash.


## Install

```go

go get github.com/jeffotoni/gcolor

```

## Use gcolor

```go

package main

import (
	// "fmt"
	. "github.com/jeffotoni/gcolor"
)

//
// start
//
func main() {

	//
	// First instantiate form
	//
	var C Color
	C.Cor = "purple"
	C.Cprintln("Purple")

	//
	// Second way to instantiate
	//
	c := new(Color)
	c.Cor = "green"
	c.Cprintln("Color Green")

	//
	// Third way to instantiate
	//
	red := Color{"red"}
	red.Cprintln("Color red")

	//
	// Fourth way to instantiate
	//
	Red.Cprintln("Red color instance otherwise")

	//
	// Yellow
	//
	Yellow.Cprintln("Yellow color!!")

}

```