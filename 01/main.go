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

	numbers := strings.Split(string(textInFile), "\n")

	increased := 0
	previews, err := strconv.Atoi(numbers[0])
	if err != nil {
		fmt.Errorf("error al convertir el numero preview %s", err)
	}
	next := numbers[1:]

	for _, n := range next {
		nextNumber, err := strconv.Atoi(n)
		if err != nil {
			fmt.Errorf("error al convertir el numero preview %s", err)
		}

		if nextNumber > previews {
			increased++
		}
		previews = nextNumber
	}

	fmt.Println(increased)

}
