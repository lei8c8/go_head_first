package main

import "fmt"

func main() {
	notes := make([]string, 7)
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"

	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(notes[6])
	fmt.Println()

	for i := 0; i < len(notes); i++ {
		fmt.Println(notes[i])
	}

	notes2 := []string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes2)

	var primes []int
	primes = append(primes, 2)
	fmt.Println(primes)
}