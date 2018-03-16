package main

import (
	"fmt"
)

func main() {
	fmt.Println(average(32, 34, 45, 76, 85, 66))
}

func average(num ...float64) float64 {
	fmt.Printf("%T", num)
	total := 0.0
	for _, v := range num {
		total += v
	}
	return total / float64(len(num))
}
