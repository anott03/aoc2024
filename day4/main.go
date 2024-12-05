package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	N := 141

	b, err := os.ReadFile("./4.in")
	if err != nil {
		log.Fatal(err)
	}

	s := string(b)
	lines := strings.Split(s, "\n")

	grid := make([][]string, N)
	for i := 0; i < N; i++ {
		grid[i] = make([]string, N)
	}

	for i := range lines {
		chars := strings.Split(lines[i], "")
		copy(grid[i], chars)
	}

	part1 := 0
	part2 := 0

	for r := 0; r < N; r++ {
		for c := 0; c < N; c++ {
			// Part 1
			if (c+3<N && grid[r][c] == "X" && grid[r][c+1] == "M" && grid[r][c+2] == "A" && grid[r][c+3] == "S"){ part1 += 1 }
			if (c+3<N && grid[r][c] == "S" && grid[r][c+1] == "A" && grid[r][c+2] == "M" && grid[r][c+3] == "X"){ part1 += 1 }
			if (r+3<N && grid[r][c] == "X" && grid[r+1][c] == "M" && grid[r+2][c] == "A" && grid[r+3][c] == "S"){ part1 += 1 }
			if (r+3<N && grid[r][c] == "S" && grid[r+1][c] == "A" && grid[r+2][c] == "M" && grid[r+3][c] == "X"){ part1 += 1 }
			if (r+3>=0 && c+3<N && grid[r][c] == "X" && grid[r+1][c+1] == "M" && grid[r+2][c+2] == "A" && grid[r+3][c+3] == "S"){ part1 += 1 }
			if (r+3>=0 && c+3<N && grid[r][c] == "S" && grid[r+1][c+1] == "A" && grid[r+2][c+2] == "M" && grid[r+3][c+3] == "X"){ part1 += 1 }
			if (r-3>=0 && c+3<N && grid[r][c] == "X" && grid[r-1][c+1] == "M" && grid[r-2][c+2] == "A" && grid[r-3][c+3] == "S"){ part1 += 1 }
			if (r-3>=0 && c+3<N && grid[r][c] == "S" && grid[r-1][c+1] == "A" && grid[r-2][c+2] == "M" && grid[r-3][c+3] == "X"){ part1 += 1 }

			// Part 2
			if r+2<N && c+2<N && grid[r][c] == "M" && grid[r+1][c+1] == "A" && grid[r+2][c+2] == "S" && grid[r+2][c] == "M" && grid[r][c+2] == "S" { part2 += 1}
			if r+2<N && c+2<N && grid[r][c] == "M" && grid[r+1][c+1] == "A" && grid[r+2][c+2] == "S" && grid[r+2][c] == "S" && grid[r][c+2] == "M" { part2 += 1}
			if r+2<N && c+2<N && grid[r][c] == "S" && grid[r+1][c+1] == "A" && grid[r+2][c+2] == "M" && grid[r+2][c] == "S" && grid[r][c+2] == "M" { part2 += 1}
			if r+2<N && c+2<N && grid[r][c] == "S" && grid[r+1][c+1] == "A" && grid[r+2][c+2] == "M" && grid[r+2][c] == "M" && grid[r][c+2] == "S" { part2 += 1}
		}
	}

	fmt.Println(part1)
	fmt.Println(part2)
}
