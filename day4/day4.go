package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	counter1 := 0
	counter2 := 0
	for scanner.Scan() {
		bothPairs := strings.Split(scanner.Text(), ",")
		firstLow, _ := strconv.Atoi(strings.Split(bothPairs[0], "-")[0])
		firstHigh, _ := strconv.Atoi(strings.Split(bothPairs[0], "-")[1])
		secondLow, _ := strconv.Atoi(strings.Split(bothPairs[1], "-")[0])
		secondHigh, _ := strconv.Atoi(strings.Split(bothPairs[1], "-")[1])

		if firstLow <= secondLow && firstHigh >= secondHigh {
			counter1++
			counter2++
		} else if secondLow <= firstLow && secondHigh >= firstHigh {
			counter1++
			counter2++
		} else if (firstHigh >= secondLow && firstHigh <= secondHigh) != (firstLow >= secondLow && firstLow <= secondHigh) {
			counter2++
		}

	}
	fmt.Println(counter1)
	fmt.Println(counter2)
}
