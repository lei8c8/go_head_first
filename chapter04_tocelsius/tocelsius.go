// tocelsius converts a temperature from F to C
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	input = strings.TrimSpace(input)
	number, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return 0, err
	}
	return number, nil
}

func main() {
	fmt.Println("Enter a temperature in F: ")
	f, err := getFloat()
	if err != nil {
		log.Fatal(err)
	}
	celsius := (f - 32) * 5 / 9
	fmt.Printf("%0.2f degrees Celsius\n", celsius)
}
