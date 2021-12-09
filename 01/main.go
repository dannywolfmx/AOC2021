package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	textInFile, err := os.ReadFile("./input.txt")

	if err != nil {
		fmt.Errorf("error al leer el archivo: %s", err)
	}

	numbers := []int{}

	text := strings.TrimSpace(string(textInFile))
	textList := strings.Split(text, "\n")
	for _, n := range textList {
		number, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		numbers = append(numbers, number)
	}

	part1(numbers)
	part2(numbers)
}

func part2(previewsList []int) {

	increased := 0

	for i := range previewsList {
		if i+4 > len(previewsList) {
			break
		}

		primeraSuma := previewsList[0+i] + previewsList[1+i] + previewsList[2+i]
		segundaSuma := previewsList[1+i] + previewsList[2+i] + previewsList[3+i]

		if segundaSuma > primeraSuma {
			increased++
		}
	}

	fmt.Println(increased)
}

func part1(numbers []int) {

	increased := 0
	previews := numbers[0]

	next := numbers[1:]

	for _, nextNumber := range next {
		if nextNumber > previews {
			increased++
		}
		previews = nextNumber
	}
	fmt.Println(increased)
}
