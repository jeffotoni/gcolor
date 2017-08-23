package gcolor

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

import "fmt"

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
