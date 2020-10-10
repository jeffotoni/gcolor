/*
* @project     gcolor
* @package     gcolor
* @author      @jeffotoni
* @size        21/08/2017
*
 */

package gcolor

import (
	"fmt"
	"strings"

	"github.com/jeffotoni/gconcat"
)

//
// Color map
//
var MapCor map[string]string

// const color
const (
	BLACK_FG  = "black"
	RED_FG    = "red"
	GREEN_FG  = "green"
	YELLOW_FG = "yellow"
	BLUE_FG   = "blue"
	PURPLE_FG = "purple"
	CYAN_FG   = "cyan"
	WHITE_FG  = "white"
)

//
// Color Struct
//
type Collor struct {
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
type ColloredPrinter interface {
	Print(Collor, string, interface{})
}

// func print color

func PrintBlack(msg ...interface{}) {
	println(BlackCor(msg))
}

func PrintRead(msg ...interface{}) {
	println(RedCor(msg))
}

func PrintCyan(msg ...interface{}) {
	println(CyanCor(msg))
}

func PrintGreen(msg ...interface{}) {
	println(GreenCor(msg))
}

func PrintYellow(msg ...interface{}) {
	println(YellowCor(msg))
}

func PrintBlue(msg ...interface{}) {
	println(BlueCor(msg))
}

func PrintPurple(msg ...interface{}) {
	println(PurpleCor(msg))
}

// print.Red()
//
//
//
func (c Collor) Print(Collor, text string, vals interface{}) {

	//
	//
	//
	var sconcat string

	//
	//
	//
	cgen := CorGeneric{Collor}
	cornow := cgen.MapColor()
	corSplit := strings.Split(cornow, "#")
	if len(corSplit) < 1 {
		return
	}

	stringt := gconcat.Build(corSplit[0], text, corSplit[1])
	sconcat = gconcat.Build(vals)

	//
	//
	//
	fmt.Println(fmt.Sprintf(stringt, sconcat))
}

//
//
//
var Printer Collor = Collor{}

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
func (c Collor) MapColor() string {
	MapCor := MapCollor()
	return MapCor[c.Cor]
}

//
// Method that builds Map Collor
//
func (c CorGeneric) MapColor() string {
	MapCor := MapCollor()
	return MapCor[c.Cor]
}

//
// Implementing the color Cprintln
//
func (c Collor) Cprintln(text string) {
	cornow := c.MapColor()
	corSplit := strings.Split(cornow, "#")
	if len(corSplit) < 1 {
		return
	}
	stringt := gconcat.Build(corSplit[0], text, corSplit[1])
	println(stringt)
}

//
// Implementing the color Cprintln to Cor Generic
//
func (c CorGeneric) Cprintln(text string) {
	cornow := c.MapColor()
	corSplit := strings.Split(cornow, "#")
	if len(corSplit) < 1 {
		return
	}
	stringt := gconcat.Build(corSplit[0], text, corSplit[1])
	fmt.Println(stringt)
}

// Sufix - Is added to the end of each color
const Sufix = "\033[0m"

// Black - Return text with black color
func BlackCor(msg ...interface{}) string {
	var sb strings.Builder
	for _, str := range msg {
		sb.WriteString(gconcat.Build(str))
	}
	return gconcat.Build("\033[0;30m", sb.String(), Sufix)
}

// Red - Return text with Red color
func RedCor(msg ...interface{}) string {
	var sb strings.Builder
	for _, str := range msg {
		sb.WriteString(gconcat.Build(str))
	}
	return gconcat.Build("\033[0;31m", sb.String(), Sufix)
}

// Green - Return text with black color
func GreenCor(msg ...interface{}) string {
	var sb strings.Builder
	for _, str := range msg {
		sb.WriteString(gconcat.Build(str))
	}
	return gconcat.Build("\033[0;32m", sb.String(), Sufix)
}

// Yellow - Return text with Yellow color
func YellowCor(msg ...interface{}) string {
	var sb strings.Builder
	for _, str := range msg {
		sb.WriteString(gconcat.Build(str))
	}
	return gconcat.Build("\033[0;33m", sb.String(), Sufix)
}

// Blue - Return text with Blue color
func BlueCor(msg ...interface{}) string {
	var sb strings.Builder
	for _, str := range msg {
		sb.WriteString(gconcat.Build(str))
	}
	return gconcat.Build("\033[0;34m", sb.String(), Sufix)
}

// Purple - Return text with Purple color
func PurpleCor(msg ...interface{}) string {
	var sb strings.Builder
	for _, str := range msg {
		sb.WriteString(gconcat.Build(str))
	}
	return gconcat.Build("\033[0;35m", sb.String(), Sufix)
}

// Cyan - Return text with Cyan color
func CyanCor(msg ...interface{}) string {
	var sb strings.Builder
	for _, str := range msg {
		sb.WriteString(gconcat.Build(str))
	}
	return gconcat.Build("\033[0;36m", sb.String(), Sufix)
}

/*
*
* @author @MarcusMann
* @size        23/08/2017
* @description Contribution made by Marcus
*
 */

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

//
// Instances of objects,
// instances of their colors
//
var Yellow CorGeneric = CorGeneric{"yellow"}
var Black CorGeneric = CorGeneric{"black"}
var Green CorGeneric = CorGeneric{"green"}
var Blue CorGeneric = CorGeneric{"blue"}
var Purple CorGeneric = CorGeneric{"purple"}
var Cyan CorGeneric = CorGeneric{"cyan"}
var Red CorGeneric = CorGeneric{"red"}
