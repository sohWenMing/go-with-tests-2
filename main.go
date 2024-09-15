package main

import (
	"fmt"
	"io"
	"os"
)

func Sum(input []int) (sum int) {
	var returned int
	for _, num := range input {
		returned += num
	}
	return returned
}

func SumAll(numbersToSum ...[]int) []int {
	var returnedSlice []int
	for _, slice := range numbersToSum {
		returnedSlice = append(returnedSlice, Sum(slice))
	}
	return returnedSlice
}

func SumAllTails(numbersToSum ...[]int) []int {
	var returnedSlice []int
	for _, slice := range numbersToSum {
		if len(slice) == 0 {
			returnedSlice = append(returnedSlice, 0)
			continue
		}
		lastIndex := len(slice) - 1
		returnedSlice = append(returnedSlice, slice[lastIndex])
	}
	return returnedSlice
}

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")

}
