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
	"strings"
	"testing"
)

func TestCorFunctions(t *testing.T) {
	tests := []struct {
		name     string
		function func(...interface{}) string
		wantSub  string
	}{
		{"BlackCor", BlackCor, "\033[0;30m"},
		{"RedCor", RedCor, "\033[0;31m"},
		{"GreenCor", GreenCor, "\033[0;32m"},
		{"YellowCor", YellowCor, "\033[0;33m"},
		{"BlueCor", BlueCor, "\033[0;34m"},
		{"PurpleCor", PurpleCor, "\033[0;35m"},
		{"CyanCor", CyanCor, "\033[0;36m"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.function("teste")
			if !strings.Contains(got, tt.wantSub) {
				t.Errorf("expected color code %s in output, got: %s", tt.wantSub, got)
			}
		})
	}
}

func TestColorStructMethods(t *testing.T) {
	c := Color{}
	if !strings.Contains(c.Red("x"), "\033[0;31m") {
		t.Error("Color.Red failed")
	}
	if !strings.Contains(c.Green("x"), "\033[0;32m") {
		t.Error("Color.Green failed")
	}
	if !strings.Contains(c.Blue("x"), "\033[0;34m") {
		t.Error("Color.Blue failed")
	}
	if !strings.Contains(c.Yellow("x"), "\033[0;33m") {
		t.Error("Color.Yellow failed")
	}
	if !strings.Contains(c.Purple("x"), "\033[0;35m") {
		t.Error("Color.Purple failed")
	}
	if !strings.Contains(c.Cyan("x"), "\033[0;36m") {
		t.Error("Color.Cyan failed")
	}
}

func TestMapColor(t *testing.T) {
	m := MapCollor()
	if m["red"] != "\033[0;31m#\033[0m" {
		t.Errorf("expected red to be \\033[0;31m#\\033[0m, got: %s", m["red"])
	}
}

func TestCollor_MapColor(t *testing.T) {
	c := Collor{Cor: "green"}
	got := c.MapColor()
	if !strings.Contains(got, "\033[0;32m") {
		t.Errorf("expected green color code, got: %s", got)
	}
}

func TestCorGeneric_MapColor(t *testing.T) {
	c := CorGeneric{Cor: "blue"}
	got := c.MapColor()
	if !strings.Contains(got, "\033[0;34m") {
		t.Errorf("expected blue color code, got: %s", got)
	}
}

func TestCorGeneric_Cprintln(t *testing.T) {
	// Test will output to stdout, but we can just call it to ensure no panic
	Cyan.Cprintln("hello")
}

func TestCollor_Cprintln(t *testing.T) {
	// Test will output to stdout, but we can just call it to ensure no panic
	c := Collor{Cor: "yellow"}
	c.Cprintln("world")
}
