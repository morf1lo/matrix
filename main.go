package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 3 {
		panic("args for matrix sizes are not provided")
	}

	// Getting matrix configuration
	rowsCount, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	columnsCount, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	// Write cooldown
	cooldownTime := time.Millisecond * time.Duration(rowsCount/columnsCount)

	w := color.New()
	w.SetWriter(os.Stdin)
	w.Add(color.FgHiGreen)

	// Creating a random matrix
	matrix := [][]int{}
	for range rowsCount {
		row := []int{}
		for range columnsCount {
			row = append(row, randomNum(0, 1))
		}
		matrix = append(matrix, row)
	}

	// Writing this matrix
	for _, row := range matrix {
		for _, col := range row {
			w.Printf("%d ", col)
			time.Sleep(cooldownTime)
		}
		w.Println()
	}
}

// Generates random number from min to max(inclusive)
func randomNum(min, max int) int {
	if max < min {
		return 0
	}

	max = max + 1

	return min + rand.Intn(max-min)
}
