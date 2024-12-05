package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func isDigit(b byte) bool {
	return '0' <= b && b <= '9'
}

func main() {
	b, err := os.ReadFile("./3.in")
	if err != nil {
		log.Fatal(err)
	}

	total := 0

	i := 0
	yay := true
	part1 := false // whether we want the answer to part 1 or part 2
	for i+7 < len(b) {
		if string(b[i:i+4]) == "do()" {
			yay = true
		} else if string(b[i:i+7]) == "don't()" {
			yay = false
		}

		if string(b[i:i+4]) == "mul(" {
			i += 4
			if !isDigit(b[i]) {
				continue
			}

			j := i
			for isDigit(b[j]) {
				j += 1
			}
			num1, err := strconv.Atoi(string(b[i:j]))
			if err != nil {
				log.Fatal(err)
			}

			if b[j] != ',' {
				continue
			}
			j += 1
			i = j

			for isDigit(b[j]) {
				j += 1
			}
			num2, err := strconv.Atoi(string(b[i:j]))
			if err != nil {
				log.Fatal(err)
			}

			if b[j] != ')' {
				continue
			}

			if part1 || yay {
				total += num1 * num2
			}
		}
		i++
	}

	fmt.Println(total)
}
