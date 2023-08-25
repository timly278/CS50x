package main

import (
	"fmt"
	"strconv"
)

/*
	AMEX
	MASTERCARD
	VISA
	INVALID
*/

const AMEX_LENGTH int = 15
const MASTERCARD_LENGTH int = 16
const VISA_LENGTH13 int = 13
const VISA_LENGTH16 int = 16

func main() {

	// 5199999999999991

	number := PromptInput()

	check := ChecksumCreditCard(int64(number))

	if check {
		fmt.Println(CheckCardType(int64(number)))
	} else {
		fmt.Println("INVALID")
	}

}

// PromptInput
func PromptInput() uint64 {

	var input uint64 = 0
	var str string
	var err error

	// make sure input is numbers
	for {
		fmt.Printf("Number: ")
		fmt.Scanln(&str)
		input, err = strconv.ParseUint(str, 10, 64)
		if err == nil {
			break
		}
		// fmt.Println("err lybatu:", err)
	}

	return input
}

// ChecksumCreditCard
func ChecksumCreditCard(num int64) bool {

	checkInput := num
	var ret bool = false
	var sumOddIndex int = 0
	var sumEvenIndex int = 0

	for num > 0 {
		sumOddIndex += int(num % 10)
		if num/10 > 0 {
			sumEvenIndex += sumOfDigit(int((num % 100) / 10 * 2))
		}

		num = num / 100
	}

	result := sumEvenIndex + sumOddIndex
	if checkInput > 0 && result%10 == 0 {
		ret = true
	}
	return ret
}

// CheckCardType include card's length and starting digits
func CheckCardType(num int64) string {
	ret := "INVALID"
	checkInput := num
	var startDigit int64 = 0
	var lengthDigit int = 0

	// check out credit card length
	for checkInput > 0 {
		lengthDigit++
		checkInput /= 10
	}

	// check out starting digits
	checkInput = num
	for checkInput > 10 {
		startDigit = checkInput % 100
		checkInput /= 10
	}
	switch startDigit {
	case 37, 34:
		if lengthDigit == AMEX_LENGTH {
			ret = "AMEX"
		}
	case 22, 55, 51, 52, 53, 54:
		if lengthDigit == MASTERCARD_LENGTH {
			ret = "MASTERCARD"
		}
	default:
		if startDigit/10 == 4 {
			if lengthDigit == VISA_LENGTH13 || lengthDigit == VISA_LENGTH16 {
				ret = "VISA"
			}
		}
	}

	return ret
}

// sumOfDigit calculates sum of digits of a number
func sumOfDigit(num int) int {
	var sum int = 0

	for num != 0 {
		sum += num % 10
		num /= 10
	}
	return sum
}
