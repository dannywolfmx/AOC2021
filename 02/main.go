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

	text := strings.TrimSpace(string(textInFile))
	textList := strings.Split(text, "\n")

	part1(textList)
	part2(textList)
}

func part2(textList []string) {
	const AID = "aid"
	const HORIZONTA = "horizontal"
	const DEPTH = "depth"
	mapa := make(map[string]uint64)

	for _, n := range textList {
		text := strings.Split(n, " ")
		num, err := strconv.ParseUint(text[1], 10, 64)
		if err != nil {
			panic("Error al convertir a numero")
		}

		switch text[0] {
		case "down":
			mapa[AID] = mapa[AID] + num
		case "up":
			mapa[AID] = mapa[AID] - num
		default:
			mapa[HORIZONTA] = mapa[HORIZONTA] + num
			mapa[DEPTH] = mapa[DEPTH] + (mapa[AID] * num)
		}
		fmt.Printf("Horizontal %v\n", mapa[HORIZONTA])
		fmt.Printf("Direccion %v\n", mapa[AID])
		fmt.Printf("Profundidad %v\n", mapa[DEPTH])
	}

	res := mapa[HORIZONTA] * mapa[DEPTH]
	fmt.Println(res)
}

func part1(textList []string) {
	mapa := make(map[string]int)

	for _, n := range textList {
		text := strings.Split(n, " ")
		num, err := strconv.Atoi(text[1])
		if err != nil {
			panic("Error al convertir a numero")
		}
		mapa[text[0]] = mapa[text[0]] + num
	}

	suma := (mapa["down"] - mapa["up"]) * mapa["forward"]
	fmt.Println(suma)
}
