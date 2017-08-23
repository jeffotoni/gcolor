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
* @author      @jeffotoni, @MarcusMann
* @size        21/08/2017
*
 */


import (
	"fmt"
	"strings"
)

//
// Color map
//
var MapCor map[string]string

const (
	RED_FG = "red"
)

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
//
//
type ColoredPrinter interface {
	Print(Color, string, ...interface{})
}

//
//
//
func (c Color) Print(Color, text string, vals interface{}) {

	//
	//
	//
	var sconcat string

	//
	//
	//
	cgen := CorGeneric{Color}
	cornow := cgen.MapColor()
	corSplit := strings.Split(cornow, "#")
	stringt := corSplit[0] + text + corSplit[1]

	//
	//
	//
	switch v := vals.(type) {

	case []string:

		for _, val := range v {

			fmt.Println("val1: ", val, "cor: ", c.Cor)
			sconcat = sconcat + val + " "
		}
	}

	//
	//
	//
	fmt.Println(fmt.Sprintf(stringt, sconcat))
}

//
//
//
var Printer Color = Color{}

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

//
// Method that builds Map color
//
func (c Color) MapColor() string {

	MapCor := MapCollor()

	return MapCor[c.Cor]
}

//
// Method that builds Map Color
//
func (c CorGeneric) MapColor() string {

	MapCor := MapCollor()

	return MapCor[c.Cor]
}

//
// Implementing the color Cprintln
//
func (c Color) Cprintln(text string) {

	cornow := c.MapColor()
	corSplit := strings.Split(cornow, "#")
	stringt := corSplit[0] + text + corSplit[1]
	fmt.Println(stringt)
}

//
// Implementing the color Cprintln to Cor Generic
//
func (c CorGeneric) Cprintln(text string) {

	cornow := c.MapColor()
	corSplit := strings.Split(cornow, "#")
	stringt := corSplit[0] + text + corSplit[1]
	fmt.Println(stringt)
}

//
// Method red, returns only string with color does not println
//
func RedCor(text string) string {

	cgen := CorGeneric{"red"}
	cornow := cgen.MapColor()
	corSplit := strings.Split(cornow, "#")
	stringt := corSplit[0] + text + corSplit[1]

	return stringt
}

//
// Method Green, returns only string with color does not println
//
func GreenCor(text string) string {

	cgen := CorGeneric{"green"}
	cornow := cgen.MapColor()
	corSplit := strings.Split(cornow, "#")
	stringt := corSplit[0] + text + corSplit[1]

	return stringt
}

//
// Method Blue, returns only string with color does not println
//
func BlueCor(text string) string {

	cgen := CorGeneric{"blue"}
	cornow := cgen.MapColor()
	corSplit := strings.Split(cornow, "#")
	stringt := corSplit[0] + text + corSplit[1]

	return stringt
}

//
// Method Yellow, returns only string with color does not println
//
func YellowCor(text string) string {

	cgen := CorGeneric{"yellow"}
	cornow := cgen.MapColor()
	corSplit := strings.Split(cornow, "#")
	stringt := corSplit[0] + text + corSplit[1]

	return stringt
}

//
// Method Purple, returns only string with color does not println
//
func PurpleCor(text string) string {

	cgen := CorGeneric{"purple"}
	cornow := cgen.MapColor()
	corSplit := strings.Split(cornow, "#")
	stringt := corSplit[0] + text + corSplit[1]

	return stringt
}

//
// Method Cyan, returns only string with color does not println
//
func CyanCor(text string) string {


// Sufix - Is added to the end of each color
const Sufix = "\033[0m"

// Color structure
type Color struct{}

// Black - Return text with black color
func (c Color) Black(msg string) string { return fmt.Sprintf("\033[0;30m%s%s", msg, Sufix) }

// Red - Return text with Red color
func (c Color) Red(msg string) string { return fmt.Sprintf("\033[0;31m%s%s", msg, Sufix) }

// Green - Return text with black color
func (c Color) Green(msg string) string { return fmt.Sprintf("\033[0;32m%s%s", msg, Sufix) }

// Yellow - Return text with Yellow color
func (c Color) Yellow(msg string) string { return fmt.Sprintf("\033[0;33m%s%s", msg, Sufix) }

// Blue - Return text with Blue color
func (c Color) Blue(msg string) string { return fmt.Sprintf("\033[0;34m%s%s", msg, Sufix) }

// Purple - Return text with Purple color
func (c Color) Purple(msg string) string { return fmt.Sprintf("\033[0;35m%s%s", msg, Sufix) }

// Cyan - Return text with Cyan color
func (c Color) Cyan(msg string) string { return fmt.Sprintf("\033[0;36m%s%s", msg, Sufix) }
