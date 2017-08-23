/*
* Go Library (C) 2017 Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
* @project     Ukkbox
* @package     main
* @author      @jeffotoni
* @size        21/08/2017
*
 */

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
