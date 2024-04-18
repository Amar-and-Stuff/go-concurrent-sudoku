package main

import (
	"fmt"
	"sync"
)

var sudoku [][]int
var waitGroup sync.WaitGroup

func main() {
	cnd := sync.NewCond(&sync.Mutex{})
	sudoku = [][]int{
		{0, 7, 0, 0, 2, 0, 0, 4, 6},
		{0, 6, 0, 0, 0, 0, 8, 9, 0},
		{2, 0, 0, 8, 0, 0, 7, 1, 5},
		{0, 8, 4, 0, 9, 7, 0, 0, 0},
		{7, 1, 0, 0, 0, 0, 0, 5, 9},
		{0, 0, 0, 1, 3, 0, 4, 8, 0},
		{6, 9, 7, 0, 0, 2, 0, 0, 8},
		{0, 5, 8, 0, 0, 0, 0, 6, 0},
		{4, 3, 0, 0, 8, 0, 0, 7, 0},
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if sudoku[i][j] == 0 {
				waitGroup.Add(1)
				go sudokuSolverRoutine(i, j, cnd)
			}
		}
	}
	waitGroup.Wait()
	fmt.Println("Solving...")
	printSudoku()
}

func sudokuSolverRoutine(row int, col int, cnd *sync.Cond) {
	state := map[int]bool{9: true, 1: true, 2: true, 3: true, 4: true, 5: true, 6: true, 7: true, 8: true}
	for true {
		cnd.L.Lock()
		checkAndRemove(state, row, col)
		if len(state) == 1 {
			sudoku[row][col] = getRemainingNumber(state)
			cnd.Broadcast()
			cnd.L.Unlock()
			break
		}
		cnd.Wait()
		cnd.L.Unlock()
	}
	waitGroup.Done()
}

func checkAndRemove(state map[int]bool, row int, col int) {
	checkHorizontal(state, row, col)
	checkVertical(state, row, col)
	checkGrid(state, row, col)
}

func checkHorizontal(state map[int]bool, row int, col int) {
	for j := 0; j < 9; j++ {
		if j != col && state[sudoku[row][j]] {
			delete(state, sudoku[row][j])
		}
	}
}

func checkVertical(state map[int]bool, row int, col int) {
	for i := 0; i < 9; i++ {
		if i != row && state[sudoku[i][col]] {
			delete(state, sudoku[i][col])
		}
	}
}

func checkGrid(state map[int]bool, row int, col int) {
	srow := (row / 3) * 3
	scol := (col / 3) * 3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if !(srow+i == row && scol+j == col) && state[sudoku[srow+i][scol+j]] {
				delete(state, sudoku[srow+i][scol+j])
			}
		}
	}
}

func getRemainingNumber(state map[int]bool) int {
	var finalElement int
	for k := range state {
		finalElement = k
		break
	}
	return finalElement
}

func printSudoku() {
	for line := range sudoku {
		fmt.Println(sudoku[line])
	}
}
