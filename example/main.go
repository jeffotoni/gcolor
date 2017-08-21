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
