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

	"github.com/jeffotoni/gcolor"
)

type Jovem struct {
	Nome string
}

func (j Jovem) Fala() string {
	return fmt.Sprintf("Meu nome é : %s", j.Nome)
}

func main() {
	j := Jovem{Nome: "Marcus Mann"}

	c := gcolor.Color{}
	cyan := c.Red(j.Fala())

	fmt.Println(cyan)
}

