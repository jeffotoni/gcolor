/*
* Go Library (C) 2017 Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
* @project     Collor
* @package     main
* @author      @jeffotoni
* @size        23/08/2017
*
 */

package gcolor

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

var old, os_r, os_w *os.File

func TestColor_Black(t *testing.T) {

	Texts1 := []string{"jeffotoni 1", "golang 2", "lagn C 3"}
	Texts2 := []int{10, 1002, 2000}
	Texts3 := "jeffotoni 2021"

	// want1 := "jeffotoni 1golang 2lagn C 3"
	// want2 := "1010022000"
	// want3 := "jeffotoni 2021"

	println(RedCor("Test Func com concat", " vamos ver red", Texts3, " ano: ", 2020))
	Print.Read("print read", 2021, " vamos ter read new now.", 700)
	Print.Yellow("print yellow", 2021, " gcolor new ", 10)
	Print.Blue("print Blue", 2021, " gcolor new ", 10)
	Print.Purple("print Purple", 2021, " gcolor new ", 10)
	Print.Black("print Black", 2021, " gcolor new ", 10)

	old = os.Stdout
	os_r, os_w, _ = os.Pipe()
	os.Stdout = os_w

	Printer.Print(YELLOW_FG, "%s", Texts1)
	Printer.Print(RED_FG, "%s", Texts2)
	Printer.Print(GREEN_FG, "%s", Texts3)

	os_w.Close()
	buf := &bytes.Buffer{}
	io.Copy(buf, os_r)
	os_r.Close()
	// Clean up.
	os.Stdout = old
	fmt.Println(buf.String())
}
