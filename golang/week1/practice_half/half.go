package main

import "fmt"

func main() {
	var bill, tax, tip float32

	fmt.Printf("Bill befor tax and tip: ")
	fmt.Scan(&bill)
	fmt.Printf("Tax percent: ")
	fmt.Scan(&tax)
	fmt.Printf("Tip percent: ")
	fmt.Scan(&tip)

	fmt.Printf("You will owe $%.2f each!\n", HalfSplit(bill, tax, tip))
}

func HalfSplit(bill float32, tax float32, tip float32) float32 {

	var half float32 = 0

	half = bill * (1 + tip/100) * (1 + tax/100) / 2

	return half
}
