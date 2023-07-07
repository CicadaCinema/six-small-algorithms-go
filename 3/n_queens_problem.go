package main

import (
	"fmt"
	"time"
)

func make_board(num_rows int) [][]string {
	result := make([][]string, num_rows)
	for i := 0; i < num_rows; i++ {
		result[i] = make([]string, num_rows)
		for j := 0; j < num_rows; j++ {
			result[i][j] = "."
		}
	}
	return result
}

func dump_board(board [][]string) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%s ", cell)
		}
		fmt.Printf("\n")
	}
}

func dump_attack_counts() {
	for _, row := range attack_counts {
		for _, cell := range row {
			fmt.Printf("%d ", cell)
		}
		fmt.Printf("\n")
	}
}

// Return true if this series of squares contains at most one queen.
func series_is_legal(board [][]string, num_rows, r0, c0, dr, dc int) bool {
	num_queens := 0
	r := r0
	c := c0

	// step through the series until we fall off the board
	for {
		// if we are off the board, break out of the loop
		if r < 0 || c < 0 || r >= num_rows || c >= num_rows {
			break
		}

		// now we know that we are still on the board

		// if this cell contains a queen, count it
		if board[r][c] == "Q" {
			num_queens += 1
		}

		// move to the next cell in the series
		r += dr
		c += dc
	}

	// returns true if a series contains at most one queen
	return num_queens <= 1
}

// Return true if the board is legal.
func board_is_legal(board [][]string, num_rows int) bool {
	for row_index := 0; row_index < num_rows; row_index++ {
		// check rows (start on the left-most cell, check in the right direction by iterating over the columns)
		// if this series is not legal, return false
		if !series_is_legal(board, num_rows, row_index, 0, 0, 1) {
			return false
		}

		// check diagonals going in the \ direction, starting from the left edge of the board
		if !series_is_legal(board, num_rows, row_index, 0, 1, 1) {
			return false
		}

		// check diagonals going in the / direction, starting from the right edge of the board
		if !series_is_legal(board, num_rows, row_index, num_rows-1, 1, -1) {
			return false
		}
	}

	for col_index := 0; col_index < num_rows; col_index++ {
		// check columns (start on the top-most cell, check in the bottom direction by iterating over the rows)
		// if this series is not legal, return false
		if !series_is_legal(board, num_rows, 0, col_index, 1, 0) {
			return false
		}

		// check diagonals going in the \ direction, starting from the top edge of the board
		if !series_is_legal(board, num_rows, 0, col_index, 1, 1) {
			return false
		}

		// check diagonals going in the / direction, starting from the top edge of the board
		if !series_is_legal(board, num_rows, 0, col_index, 1, -1) {
			return false
		}

	}

	// we have checked all rows, columns and diagonals
	return true
}

// Return true if the board is legal and a solution.
func board_is_a_solution(board [][]string, num_rows int) bool {
	num_queens := 0
	for _, row := range board {
		for _, cell := range row {
			if cell == "Q" {
				num_queens += 1
			}
		}
	}

	// solution && legal
	return num_queens == num_rows && board_is_legal(board, num_rows)
}

// Try placing a queen at position [r][c].
// Return true if we find a legal board.
func place_queens_1(board [][]string, num_rows, r, c int) bool {
	if r >= num_rows {
		// we have fallen off the board, so check this assignment of queens
		return board_is_a_solution(board, num_rows)
	}

	// find the next square's location
	next_c := (c + 1) % num_rows
	// if c is at most num_rows-2, then c+1 is at most num_rows-1, then (c+1)/num_rows is 0 and the row index is not incremented
	// otherwise, the result of this floor division is 1 and the row index is incremented
	next_r := r + ((c + 1) / num_rows)

	// see what happens if we do not place a queen in square (r, c)
	if place_queens_1(board, num_rows, next_r, next_c) {
		// if this returned a solution, then we have found a solution
		return true
	}

	// now we know that leaving (r, c) without a queen cannot return a solution

	// see what happens if we try putting a queen there
	board[r][c] = "Q"
	if place_queens_1(board, num_rows, next_r, next_c) {
		// if this returned a solution, then we have found a solution
		return true
	}

	// now we know that neither putting a queen, nor leaving this cell as-is works
	// there is no possible solution from the board position that we were given
	// so reset this cell and return false to indicate that there is no possible solution
	board[r][c] = "."
	return false
}

func place_queens_2(board [][]string, num_rows, r, c, num_placed int) bool {
	if num_placed == num_rows {
		// we have already placed all our queens, so check this assignment of queens
		return board_is_a_solution(board, num_rows)
	}

	if r >= num_rows {
		// we have no yet placed all the queens, but we have fallen off the board
		// so there is no solution
		return false
	}

	// find the next square's location
	next_c := (c + 1) % num_rows
	// if c is at most num_rows-2, then c+1 is at most num_rows-1, then (c+1)/num_rows is 0 and the row index is not incremented
	// otherwise, the result of this floor division is 1 and the row index is incremented
	next_r := r + ((c + 1) / num_rows)

	// see what happens if we do not place a queen in square (r, c)
	if place_queens_2(board, num_rows, next_r, next_c, num_placed) {
		// if this returned a solution, then we have found a solution
		return true
	}

	// now we know that leaving (r, c) without a queen cannot return a solution

	// see what happens if we try putting a queen there
	board[r][c] = "Q"
	if place_queens_2(board, num_rows, next_r, next_c, num_placed+1) {
		// if this returned a solution, then we have found a solution
		return true
	}

	// now we know that neither putting a queen, nor leaving this cell as-is works
	// there is no possible solution from the board position that we were given
	// so reset this cell and return false to indicate that there is no possible solution
	board[r][c] = "."
	return false
}

func place_queens_3(board [][]string, num_rows, r, c, num_placed int) bool {
	if num_placed == num_rows {
		// we have already placed all our queens, so check this assignment of queens
		return board_is_a_solution(board, num_rows)
	}

	if r >= num_rows {
		// we have no yet placed all the queens, but we have fallen off the board
		// so there is no solution
		return false
	}

	// find the next square's location
	next_c := (c + 1) % num_rows
	// if c is at most num_rows-2, then c+1 is at most num_rows-1, then (c+1)/num_rows is 0 and the row index is not incremented
	// otherwise, the result of this floor division is 1 and the row index is incremented
	next_r := r + ((c + 1) / num_rows)

	// see what happens if we do not place a queen in square (r, c)
	if place_queens_3(board, num_rows, next_r, next_c, num_placed) {
		// if this returned a solution, then we have found a solution
		return true
	}

	// now we know that leaving (r, c) without a queen cannot return a solution

	// only place a queen if the attack count is 0 for the test cell
	if attack_counts[r][c] == 0 {
		// see what happens if we try putting a queen there
		board[r][c] = "Q"
		adjust_attack_counts(num_rows, r, c, 1)
		if place_queens_3(board, num_rows, next_r, next_c, num_placed+1) {
			// if this returned a solution, then we have found a solution
			return true
		}
		// we did not return true, so need to reset
		board[r][c] = "."
		adjust_attack_counts(num_rows, r, c, -1)
	}

	return false
}

var attack_counts [][]int

func make_attack_counts(num_rows int) [][]int {
	result := make([][]int, num_rows)
	for i := 0; i < num_rows; i++ {
		result[i] = make([]int, num_rows)
		for j := 0; j < num_rows; j++ {
			result[i][j] = 0
		}
	}
	return result
}

// adjust the attack counts for all the cells which are hit by a queen at (r,c) by adding val to the cell value
func adjust_attack_counts(num_rows, r, c, val int) {
	// adjust rows and columns
	for i := 0; i < num_rows; i++ {
		attack_counts[i][c] += val
		attack_counts[r][i] += val

	}
	// adjust diagonals
	for i := 1; i < num_rows; i++ {
		set_r := r + i
		set_c := c + i
		if set_r < num_rows && set_c < num_rows {
			attack_counts[set_r][set_c] += val
		}
		set_r = r - i
		set_c = c - i
		if set_r >= 0 && set_c >= 0 {
			attack_counts[set_r][set_c] += val
		}
		set_r = r - i
		set_c = c + i
		if set_r >= 0 && set_c < num_rows {
			attack_counts[set_r][set_c] += val
		}
		set_r = r + i
		set_c = c - i
		if set_r < num_rows && set_c >= 0 {
			attack_counts[set_r][set_c] += val
		}
	}
	// in fact we counted (r,c) twice
	attack_counts[r][c] -= val

}

func main() {
	const num_rows = 13
	board := make_board(num_rows)
	attack_counts = make_attack_counts(num_rows)

	start := time.Now()
	//success := place_queens_1(board, num_rows, 0, 0)
	//success := place_queens_2(board, num_rows, 0, 0, 0)
	success := place_queens_3(board, num_rows, 0, 0, 0)

	elapsed := time.Since(start)
	if success {
		fmt.Println("Success!")
		dump_board(board)
	} else {
		fmt.Println("No solution")
	}
	fmt.Printf("Elapsed: %f seconds\n", elapsed.Seconds())
}
