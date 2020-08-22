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
* @author      @marcus
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

// func TestColor_Cyan(t *testing.T) {

// 	type args struct {
// 		msg string
// 	}
// 	tests := []struct {
// 		name string
// 		c    Color
// 		args args
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 	}

// 	for _, tt := range tests {

// 		t.Run(tt.name, func(t *testing.T) {

// 			c := Color{}
// 			if got := c.Cyan(tt.args.msg); got != tt.want {
// 				t.Errorf("Color.Cyan() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestColor_Purple(t *testing.T) {

// 	type args struct {
// 		msg string
// 	}
// 	tests := []struct {
// 		name string
// 		c    Color
// 		args args
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := Color{}
// 			if got := c.Purple(tt.args.msg); got != tt.want {
// 				t.Errorf("Color.Purple() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestColor_Blue(t *testing.T) {

// 	type args struct {
// 		msg string
// 	}
// 	tests := []struct {
// 		name string
// 		c    Color
// 		args args
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := Color{}
// 			if got := c.Blue(tt.args.msg); got != tt.want {
// 				t.Errorf("Color.Blue() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestColor_Yellow(t *testing.T) {

// 	type args struct {
// 		msg string
// 	}
// 	tests := []struct {
// 		name string
// 		c    Color
// 		args args
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := Color{}
// 			if got := c.Yellow(tt.args.msg); got != tt.want {
// 				t.Errorf("Color.Yellow() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestColor_Green(t *testing.T) {

// 	type args struct {
// 		msg string
// 	}
// 	tests := []struct {
// 		name string
// 		c    Color
// 		args args
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := Color{}
// 			if got := c.Green(tt.args.msg); got != tt.want {
// 				t.Errorf("Color.Green() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestColor_Red(t *testing.T) {

// 	type args struct {
// 		msg string
// 	}
// 	tests := []struct {
// 		name string
// 		c    Color
// 		args args
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			c := Color{}
// 			if got := c.Red(tt.args.msg); got != tt.want {
// 				t.Errorf("Color.Red() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
var old, os_r, os_w *os.File

func TestColor_Black(t *testing.T) {

	Texts1 := []string{"jeffotoni 1", "golang 2", "lagn C 3"}
	Texts2 := []int{10, 1002, 2000}
	Texts3 := "jeffotoni 2021"

	// want1 := "jeffotoni 1golang 2lagn C 3"
	// want2 := "1010022000"
	// want3 := "jeffotoni 2021"

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
