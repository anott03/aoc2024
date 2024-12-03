package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	f, err := os.Open("./2.in")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		nums := make([]int, len(parts))
		for i := range parts {
			n, err := strconv.Atoi(parts[i])
			if err != nil {
				log.Fatal(err)
			}
			nums[i] = n
		}

		yayay := false

		for k := 0; k < len(nums); k++ {
			idxs := make([]int, len(nums)-1)
			idx := 0
			yay := true
			for i := range idxs {
				idxs[i] = idx
				idx += 1
				if idx == k {
					idx += 1
				}
			}

			numss := make([]int, len(nums)-1)
			for i := range(numss) {
				numss[i] = nums[idxs[i]]
			}

			sorted := slices.IsSorted(numss)
			if !sorted {
				for i, j := 0, len(numss)-1; i < j; i, j = i+1, j-1 {
					numss[i], numss[j] = numss[j], numss[i]
				}
				sorted = slices.IsSorted(numss)
			}

			for i := 1; i < len(numss); i++ {
				diff := abs(numss[i] - numss[i-1])
				if 1 > diff || 3 < diff {
					yay = false
					break
				}
			}

			yayay = sorted && yay
		}

		if yayay {
			count += 1
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count) // correct answre for part 2 is 255
}
