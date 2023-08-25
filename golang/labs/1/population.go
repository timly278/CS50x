package main

import (
	"fmt"
	// "go/doc"
)

const MIN_SIZE int = 9

func main() {
	// TODO: Prompt for start size
	endSize := 0
	startSize := PromptStartSize()

	// TODO: Prompt for end size
	for endSize <= startSize {
		endSize = PromptEndSize()
	}

	fmt.Println("StartSize is", startSize)
	fmt.Println("EndSize is", endSize)

	// TODO: Calculate number of years until we reach threshold
	Years := RequiredYears(startSize, endSize)

	// TODO: Print number of years
	fmt.Println("Years:", Years)

}

// RequiredYears to calculate the required years for the llama population to reach that end size
func RequiredYears(start, end int) int {

	var numOfYears int = 0
	sizeLlamas := start

	for sizeLlamas < end {
		sizeLlamas = sizeLlamas + sizeLlamas/3 - sizeLlamas/4
		numOfYears++
	}
	return numOfYears
}

// PromptStartSize to prompt user enter the start size of llama
func PromptStartSize() int {

	size := 0

	for size < MIN_SIZE {
		fmt.Printf("Start size: ")
		fmt.Scan(&size)
	}

	return size
}

// PromptEndSize to prompt user enter the start size of llama
func PromptEndSize() int {

	var size int

	fmt.Printf("End size: ")
	fmt.Scan(&size)

	return size
}
