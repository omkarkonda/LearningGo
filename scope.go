package main

import (
	"fmt"
)

func omk() {
	x := 10
	func1 := func() int {
		x++
		return x
	}
	fmt.Println(func1())

	wrapper := wrapperFunc()
	wrapper()
}

func wrapperFunc() func() int {
	x := 6

	return func() int {
		return x
	}
}
