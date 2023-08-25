package main

import "fmt"

const MAX_HEIGHT int = 8
const MIN_HEIGHT int = 1

func main() {

	GenPyramid(1)
	fmt.Printf("\n")
	GenPyramid(7)
	fmt.Printf("\n")
	GenPyramid(6)
	fmt.Printf("\n")
	GenPyramid(5)
	fmt.Printf("\n")

	height := PromptHeight()
	GenPyramid(height)
	fmt.Println("Thank you for playing!")
}

// PromptEndSize to prompt user enter the start size of llama
func PromptHeight() int {

	var height int = 0

	for height < MIN_HEIGHT || height > MAX_HEIGHT {
		fmt.Printf("Height: ")
		fmt.Scan(&height)
	}

	return height
}

// GenPyramid to build the pyramid with [1,8] of height
func GenPyramid(height int) {

	if height < MIN_HEIGHT || height > MAX_HEIGHT {
		return
	} else {
		center := height
		line := make([]string, height*2+2)

		// Initiate slice
		for i := range line {
			line[i] = " "
		}

		// Create the pyramid
		for i := 1; i <= height; i++ {
			line[center+i+1] = "#"
			line[center-i] = "#"

			// Print out the line
			for _, v := range line {
				fmt.Printf("%s", v)
			}
			fmt.Printf("\n")
		}
	}
}
