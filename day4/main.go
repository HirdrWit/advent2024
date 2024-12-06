package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

var (
	data   = readData("./day4/data.txt")
	sample = [][]string{
		{".", ".", ".", ".", "X", "X", "M", "A", "S", "."},
		{".", "S", "A", "M", "X", "M", "S", ".", ".", "."},
		{".", ".", ".", "S", ".", ".", "A", ".", ".", "."},
		{".", ".", "A", ".", "A", ".", "M", "S", ".", "X"},
		{"X", "M", "A", "S", "A", "M", "X", ".", "M", "M"},
		{"X", ".", ".", ".", ".", ".", "X", "A", ".", "A"},
		{"S", ".", "S", ".", "S", ".", "S", ".", "S", "S"},
		{".", "A", ".", "A", ".", "A", ".", "A", ".", "A"},
		{".", ".", "M", ".", "M", ".", "M", ".", "M", "M"},
		{".", "X", ".", "X", ".", "X", "M", "A", "S", "X"},
	}

	sample2 = [][]string{
		{".", "M", ".", "S", ".", ".", ".", ".", ".", "."},
		{".", ".", "A", ".", ".", "M", "S", "M", "S", "."},
		{".", "M", ".", "S", ".", "M", "A", "A", ".", "."},
		{".", ".", "A", ".", "A", "S", "M", "S", "M", "."},
		{".", "M", ".", "S", ".", "M", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
		{"S", ".", "S", ".", "S", ".", "S", ".", "S", "."},
		{".", "A", ".", "A", ".", "A", ".", "A", ".", "A"},
		{"M", ".", "M", ".", "M", ".", "M", ".", "M", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", ".", "."},
	}
)

func main() {
	answer1(data)
	answer2(data)
}

func answer2(data [][]string) {
	x_mas := "MAS"
	count := 0
	for y := 1; y < len(data)-1; y++ {
		for x := 1; x < len(data[y])-1; x++ {
			if data[y][x] == "A" {
				word1 := strings.Join([]string{data[y-1][x-1], data[y][x], data[y+1][x+1]}, "")
				word2 := strings.Join([]string{data[y-1][x+1], data[y][x], data[y+1][x-1]}, "")
				if (word1 == x_mas || reverse(word1) == x_mas) && (word2 == x_mas || reverse(word2) == x_mas) {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func answer1(data [][]string) {
	xmas := "XMAS"
	count := 0
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			// check left
			if x-3 > 0 {
				word := strings.Join(data[y][x-4:x], "")
				if word == xmas || reverse(word) == xmas {
					count++
				}
			}
			// check up
			if y-3 >= 0 {
				word := strings.Join([]string{data[y-3][x], data[y-2][x], data[y-1][x], data[y][x]}, "")
				if word == xmas || reverse(word) == xmas {
					count++
				}
			}

			// check back up and left
			if x-3 >= 0 && y-3 >= 0 {
				word := strings.Join([]string{data[y-3][x-3], data[y-2][x-2], data[y-1][x-1], data[y][x]}, "")
				if word == xmas || reverse(word) == xmas {
					count++
				}
			}

			// check forward up and right
			if x+3 < len(data) && y-3 >= 0 {
				word := strings.Join([]string{data[y-3][x+3], data[y-2][x+2], data[y-1][x+1], data[y][x]}, "")
				if word == xmas || reverse(word) == xmas {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func reverse(s string) string {
	var reversed string = ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}
	return reversed
}

func readData(filePath string) [][]string {
	result := make([][]string, 0)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := make([]string, 0)
		for _, ch := range scanner.Text() {
			row = append(row, string(ch))
		}
		result = append(result, row)
	}
	if err = scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return result
}
