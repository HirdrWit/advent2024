package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var (
	data = readData("./day3/data.txt")
	re   = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
)

func main() {
	answer1()
	answer2()
}

func answer1() {
	fmt.Println(findResult(data))
}

func answer2() {
	cleanRE := regexp.MustCompile(`don't|do|mul\(\d{1,3},\d{1,3}\)`)

	cleanedData := make([]string, 0)
	for _, line := range data {
		cleanedData = append(cleanedData, cleanRE.FindAllString(line, -1)...)
	}

	do := make([]string, 0)
	doActive := true
	for _, line := range cleanedData {
		switch line {
		case "don't":
			doActive = false
		case "do":
			doActive = true
		default:
			if doActive {
				do = append(do, line)
			}
		}
	}
	fmt.Println(findResult(do))
}

func findResult(d []string) int {
	result := 0
	for _, line := range d {
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			i, _ := strconv.Atoi(match[1])
			j, _ := strconv.Atoi(match[2])
			result += i * j
		}
	}
	return result
}

func readData(filePath string) []string {
	result := make([]string, 0)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
