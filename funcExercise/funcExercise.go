package main

import "fmt"

func main() {
	fmt.Println(ex1(2))
}

func ex1(num int) (int, bool) {
	return num / 2, num%2 == 0
}
