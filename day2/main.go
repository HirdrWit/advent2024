package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

var matrix = readCsvFile("data.csv")

func main() {
	answer1()
	answer2()
}

func answer1() {
	safeCount := 0
	for _, row := range matrix {
		if isSafe(row) {
			safeCount++
		}
	}
	fmt.Println(safeCount)
}

func answer2() {
	safeCount := 0
	for _, row := range matrix {
		if isSafe(row) {
			safeCount++
			continue
		}
		for i := 0; i < len(row); i++ {
			if isSafe(makeSkipArray(row, i)) {
				safeCount++
				break
			}
		}
	}
	fmt.Println(safeCount)
}

func makeSkipArray(row []int, skip int) []int {
	result := make([]int, 0, len(row)-1)
	for i, v := range row {
		if i == skip {
			continue
		}
		result = append(result, v)
	}
	return result
}

func isSafe(row []int) bool {
	left, right := 0, 1
	increasing := row[0] < row[len(row)-1]

	for right < len(row) {
		small, big := row[right], row[left]
		if increasing {
			small, big = row[left], row[right]
		}
		if small >= big || (big-small) > 3 {
			return false
		}
		left = right
		right++
	}
	return true
}

func readCsvFile(filePath string) [][]int {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath+" ", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.FieldsPerRecord = -1 // Allow variable number of fields
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	result := make([][]int, 0)
	for _, record := range records {
		row := make([]int, 0, len(record))
		for _, cell := range record {
			v, _ := strconv.Atoi(cell)
			row = append(row, v)
		}
		result = append(result, row)
	}

	return result
}
