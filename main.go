package main

func Sum(input []int) (sum int) {
	var returned int
	for _, num := range input {
		returned += num
	}
	return returned
}
