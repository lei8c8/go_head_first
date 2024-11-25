package main

import "fmt"

func double(number *int) {
	*number = *number * 2
}

func main() {
	var num int = 6
	double(&num)
	fmt.Println(num)
}