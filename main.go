package main

import "fmt"

func main() {
	fmt.Println("hello world")
	fmt.Printf("%#X - %b - %o \n", 42, 42, 42)
	fmt.Println("-----------------------------------")

	for i := 0; i < 200; i++ {
		fmt.Printf(" %d \t %b - %#x \t %U \n", i, i, i, i)
	}
}
