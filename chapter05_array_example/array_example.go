package main

import "fmt"

func main() {
	notes := [7]string{"do", "re", "mi", "fa", "so", "la", "ti"}
	index := 1
	fmt.Println(index, notes[index])
	index = 3
	fmt.Println(index, notes[index])
	fmt.Println()

	for i := 0; i <= 2; i++ {
		fmt.Println(i, notes[i])
	}

	for i := 0; i < len(notes); i++ {
		fmt.Println(i, notes[i])
	}

	fmt.Println()

	for index, note := range notes {
		fmt.Println(index, note)
	}
}
