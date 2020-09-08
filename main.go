package main

import "fmt"

func superAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println(i, numbers[i])
	// }

	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6, 1, 2, 3, 4, 5)
	fmt.Println(result)
}
