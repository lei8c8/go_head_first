package main 

import "fmt"

func main() {
	numbers := [3]float64{71.8, 56.2, 89.5}
	var sum float64 = 0
	for _, v := range numbers {
		sum += v
	}
	fmt.Println("Sum is:", sum)
}