# gocolor

A simple lib to help you color when developing some bash applications using ANSI-COLORS.

The goal is to make the lib leaner and leaner.

The gcolor was made just with this intuition to leave the texts in the colored bash.

You can collaborate by giving a "Fork" and sending us a "pull request", lib gcolor was just an academic play with only didactic input, which we use in our projects, something fun to do and learn.


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
	Printer.Print(RED_FG, "test string %s", Texts)

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