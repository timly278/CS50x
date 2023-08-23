package main

import "fmt"

func main() {
	var max, min int

	fmt.Printf("Minimum: ")
	fmt.Scan(&min)
	fmt.Printf("Maximum: ")
	fmt.Scan(&max)

	prime := PrimeNumber(min, max)
	fmt.Println(prime)

}

func PrimeNumber(min, max int) []int {
	var primeSlice []int
	j := 0

	for i := min; i <= max; i++ {
		j = 2
		for ; j < i; j++ {
			if i%j == 0 {
				break
			}
		}
		if j == i {
			primeSlice = append(primeSlice, i)
		}
	}

	return primeSlice
}
