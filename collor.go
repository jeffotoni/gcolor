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

package gcolor

import (
	"fmt"
	"strings"
)

//
// Color map
//
var MapCor map[string]string

//
// Color Struct
//
type Color struct {
	Cor string
}

//
// version Generic
//
type CorGeneric struct {
	Cor string
}

//
// Implementing the color map
//
func MapCollor() map[string]string {

	MapCor = make(map[string]string)

	MapCor["black"] = "\033[0;30m#\033[0m"
	MapCor["red"] = "\033[0;31m#\033[0m"
	MapCor["green"] = "\033[0;32m#\033[0m"
	MapCor["yellow"] = "\033[0;33m#\033[0m"
	MapCor["blue"] = "\033[0;34m#\033[0m"
	MapCor["purple"] = "\033[0;35m#\033[0m"
	MapCor["cyan"] = "\033[0;36m#\033[0m"
	MapCor["white"] = "\033[37m#\033[0m"

	return MapCor
}

func (c Color) MapColor() string {

	MapCor := MapCollor()

	return MapCor[c.Cor]
}

func (c Color) Cprintln(text string) {

	cornow := c.MapColor()
	corSplit := strings.Split(cornow, "#")
	stringt := corSplit[0] + text + corSplit[1]
	fmt.Println(stringt)
}

func (c CorGeneric) MapColor() string {

	MapCor := MapCollor()

	return MapCor[c.Cor]
}

func (c CorGeneric) Cprintln(text string) {

	cornow := c.MapColor()
	corSplit := strings.Split(cornow, "#")
	stringt := corSplit[0] + text + corSplit[1]
	fmt.Println(stringt)
}

var Yellow CorGeneric = CorGeneric{"yellow"}
var Black CorGeneric = CorGeneric{"black"}
var Green CorGeneric = CorGeneric{"green"}
var Blue CorGeneric = CorGeneric{"blue"}
var Purple CorGeneric = CorGeneric{"purple"}
var Cyan CorGeneric = CorGeneric{"cyan"}
var Red CorGeneric = CorGeneric{"red"}
