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

	rowLen, err0 := strconv.Atoi(os.Args[1])
	columnsCount, err1 := strconv.Atoi(os.Args[2])
	if err0 != nil || err1 != nil {
		panic(err0)
	}

	cooldownTime := time.Millisecond * time.Duration(rowLen/columnsCount)

	w := color.New()
	w.SetWriter(os.Stdin)
	w.Add(color.FgHiGreen)

	matrix := [][]int{}
	for range rowLen {
		row := []int{}
		for range columnsCount {
			row = append(row, randomNum(0, 1))
		}
		matrix = append(matrix, row)
	}

	for _, row := range matrix {
		for _, col := range row {
			w.Printf("%d ", col)
			time.Sleep(cooldownTime)
		}
		w.Println()
	}
}

func randomNum(min, max int) int {
	if max < min {
		return 0
	}

	max = max + 1

	return min + rand.Intn(max-min)
}
