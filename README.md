# gocolor

A simple lib to help you color when developing some bash applications using ANSI-COLORS.

The goal is to make the lib leaner and leaner.

The gcolor was made just with this intuition to leave the texts in the colored bash.

You can collaborate by giving a "Fork" and sending us a "pull request", lib gcolor was just an academic play with only didactic input, which we use in our projects, something fun to do and learn.


## Install

```go

go get github.com/jeffotoni/gcolor

```

## Use gcolor

```go

package main

import (
	// "fmt"
	"github.com/jeffotoni/gcolor"
)

// How to get a color
package main

import (
	"fmt"

	"github.com/jeffotoni/gcolor"
)

func main() {
	c := gcolor.Color{}

	fmt.Println(c.Cyan("Hi!"))
}
```