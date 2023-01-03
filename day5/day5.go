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

	var boxes []string = make([]string, 9)

	parseBoxes := true
	for scanner.Scan() {
		line := scanner.Text()

		if parseBoxes {
			parseBoxes = !strings.Contains(line, "1")
			if !parseBoxes {
				continue
			}
			lineWidth := 9 * 4
			for i := 0; i < lineWidth; i += 4 {
				box := line[i : i+3]
				if box != "   " {
					boxes[i/4] = box + boxes[i/4]
					if strings.Contains(box, "1") {
						parseBoxes = false
					}
				}

			}

		} else {
			if line == "" {
				continue
			}
			field := strings.Fields(line)
			moves, _ := strconv.ParseInt(field[1], 10, 32)
			from, _ := strconv.ParseInt(field[3], 10, 32)
			to, _ := strconv.ParseInt(field[5], 10, 32)

			from--
			to--
			// boxes[from] = strings.TrimSuffix(boxes[from], box)

			box := boxes[from][len(boxes[from])-int(moves)*3:]
			boxes[from] = strings.TrimSuffix(boxes[from], box)
			boxes[to] += box

			// for i := len(box); i > 0; i -= 3 {
			// 	b := box[i-3 : i]
			// 	boxesOne[to] += b
			// }

		}

	}

	for _, box := range boxes {
		fmt.Printf("%s, ", box[len(box)-3:])
	}
}
