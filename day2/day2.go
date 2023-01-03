package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {

	file, _ := os.Open("input.txt")
	data, _ := ioutil.ReadAll(file)
	splitInput := strings.Split(string(data[:]), "\n")

	// A rock X
	// B paper Y
	// C scissor Z
	score1 := map[string]map[string]int{
		"A": {
			"X": 1 + 3,
			"Y": 2 + 6,
			"Z": 3 + 0,
		},
		"B": {
			"X": 1 + 0,
			"Y": 2 + 3,
			"Z": 3 + 6,
		},
		"C": {
			"X": 1 + 6,
			"Y": 2 + 0,
			"Z": 3 + 3,
		},
	}

	// X loss
	// Y draw
	// Z win

	// A rock
	// B paper
	// C scissor
	score2 := map[string]map[string]int{
		"A": {
			"X": 3 + 0,
			"Y": 1 + 3,
			"Z": 2 + 6,
		},
		"B": {
			"X": 1 + 0,
			"Y": 2 + 3,
			"Z": 3 + 6,
		},
		"C": {
			"X": 2 + 0,
			"Y": 3 + 3,
			"Z": 1 + 6,
		},
	}

	var cumSum1 int
	var cumSum2 int

	for _, s := range splitInput {
		var moves = strings.Split(s, " ")
		cumSum1 += score1[moves[0]][moves[1]]
		cumSum2 += score2[moves[0]][moves[1]]
	}

	fmt.Println(cumSum1)
	fmt.Println(cumSum2)
}
