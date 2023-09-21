package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	var dataFile, dnaFile *os.File

	// Check command-line arguments
	lenArgs := len(args)
	if lenArgs < 2 || lenArgs > 3 {
		fmt.Println("lack of argument!")
		return
	}

	// Open files
	dataFile, dnaFile = openFiles(args[1], args[2])
	if dataFile == nil {
		fmt.Println("can't open file!")
		return
	}

	defer dataFile.Close()
	defer dnaFile.Close()

	// Read files
	dataReader, dnaReader, err := readFiles(dataFile, dnaFile)
	if err != nil {
		fmt.Println("can't read file!")
		return
	}

	fmt.Println(findMatchingPerson(dataReader, string(dnaReader)))

}

// findMatchingPerson compare formed data with database and return name of person who matches with the given DNA
func findMatchingPerson(database [][]string, DNA string) string {

	STRs := formStrDataStruct(database[0], DNA)
	var result string

	for _, data := range database[1:] {
		if isEqualStringSlice(data[1:], STRs) {
			result = data[0]
			return result
		}
	}
	result = "No match"

	return result
}

// isEqualStringSlice Compare two slice of strings return true if equal
func isEqualStringSlice(slice1, slice2 []string) bool {

	if len(slice1) != len(slice2) {
		return false
	}

	for i, val1 := range slice1 {
		if slice2[i] != val1 {
			return false
		}
	}

	return true
}

// formStrDataStruct form STR data not include name field
func formStrDataStruct(dataForm []string, DNA string) []string {
	var strStruct []string

	for _, data := range dataForm[1:] {
		number := longestMatch(data, DNA)
		strStruct = append(strStruct, strconv.Itoa(number))
	}

	return strStruct
}

// isEqualString compare two strings
func isEqualString(str1, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}

	for i := range str1 {
		if str2[i] != str1[i] {
			return false
		}
	}

	return true
}

// longestMatch return the longest times the STR appear in dna consecutively
func longestMatch(STR string, DNA string) int {
	count, result := 0, 0
	lenSTR := len(STR)

	i := 0
	for i < len(DNA)-lenSTR {

		if isEqualString(DNA[i:i+lenSTR], STR) {
			count++
			i += lenSTR
		} else {
			i++
			if count > result {
				result = count
			}
			count = 0
		}
	}

	if count > result {
		result = count
	}
	return result
}

// openFiles takes args as inputs and returns dataFile, dnaFile
func openFiles(dataFileName, dnaFileName string) (*os.File, *os.File) {

	dataDes, err := os.Open(dataFileName)
	if err != nil {
		return nil, nil
	}

	dnaDes, err := os.Open(dnaFileName)
	if err != nil {
		return nil, nil
	}

	return dataDes, dnaDes
}

// readFiles read .csv and .text files respectively
func readFiles(dataFile, dnaFile *os.File) ([][]string, []byte, error) {

	dataReader, err := csv.NewReader(dataFile).ReadAll()
	if err != nil {
		return nil, nil, err
	}

	// dnaReader, err := bufio.NewReader(dnaFile).ReadBytes('\n')
	dnaReader, err := io.ReadAll(dnaFile)
	if err != nil {
		return nil, nil, err
	}

	return dataReader, dnaReader, nil
}
