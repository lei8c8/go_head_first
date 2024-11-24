package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("Hello, Go!", "Hello Python")
	fmt.Println(math.Floor(2.75))
	fmt.Println(strings.Title("\"head first go\"\\"))
	fmt.Println('A')

	var myInt int
	var myFloat float64
	fmt.Println(myInt, myFloat)
}
