package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

var array1, array2 = readCsvFile("data.csv")

func main() {
	answer1()
	answer2()
}

func answer1() {
	sort.Ints(array1)
	sort.Ints(array2)
	result := 0
	for i := 0; i < len(array1); i++ {
		if array1[i] > array2[i] {
			result += array1[i] - array2[i]
			continue
		}
		result += array2[i] - array1[i]
	}
	fmt.Println(result)
}

func answer2() {
	m := make(map[int]int)
	for _, v := range array2 {
		m[v] = m[v] + 1
	}

	sum := 0
	for _, v := range array1 {
		if _, ok := m[v]; ok {
			sum += v * m[v]
		}
	}
	fmt.Println(sum)
}

func readCsvFile(filePath string) ([]int, []int) {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}
	array1 := make([]int, len(records))
	array2 := make([]int, len(records))
	for i, record := range records {
		array1[i], _ = strconv.Atoi(record[0])
		array2[i], _ = strconv.Atoi(record[1])
	}

	return array1, array2
}
