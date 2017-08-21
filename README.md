# gocolor

A simple lib to help you color when developing some bash application using ANSI-COLORS.

The goal is to make the lib leaner and leaner.

The gocollor was made only with this intúito of leaving its characters in the colored bash.


## Install

```go

go get github.com/jeffotoni/gocollor

```

## Use gocollor

```go

package main

import (
	// "fmt"
	. "github.com/jeffotoni/gocollor"
)

//
// start
//
func main() {

	//
	// First instantiate form
	//
	var color Collor
	color.Cor = "purple"
	color.Cprintln("Purple")

	//
	// Second way to instantiate
	//
	c := new(Collor)
	c.Cor = "green"
	c.Cprintln("Start port")

	//
	// Third way to instantiate
	//
	red := Collor{"red"}
	red.Cprintln("Start port red")

	//
	// Fourth way to instantiate
	//
	Red.Cprintln("Collor red2")

}

```