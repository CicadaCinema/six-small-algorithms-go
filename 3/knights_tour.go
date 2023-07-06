package main

import (
	"fmt"
	"time"
)

// The board dimensions.
const num_rows = 8
const num_cols = num_rows

// Whether we want an open or closed tour.
const require_closed_tour = false

// Value to represent a square that we have not visited.
const unvisited = -1

// Define offsets for the knight's movement.
type Offset struct {
	dr, dc int
}

var move_offsets []Offset

var num_calls int64

func initialize_offsets() {
	// SLOW
	/*
	move_offsets = []Offset{
		Offset{+2, -1},
		Offset{+2, +1},
		Offset{-2, -1},
		Offset{-2, +1},
		Offset{-1, +2},
		Offset{+1, +2},
		Offset{-1, -2},
		Offset{+1, -2},
	}
	*/

	move_offsets = []Offset{
		Offset{-2, -1},
		Offset{-1, -2},
		Offset{+2, -1},
		Offset{+1, -2},
		Offset{-2, +1},
		Offset{-1, +2},
		Offset{+2, +1},
		Offset{+1, +2},
	}

}

func make_board(num_rows int, num_cols int) [][]int {
	result := make([][]int, num_rows)
	for i := 0; i < num_rows; i++ {
		result[i] = make([]int, num_cols)
		for j := 0; j < num_rows; j++ {
			result[i][j] = unvisited
		}
	}
	return result
}

func dump_board(board [][]int) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Printf("%02d ", cell)
		}
		fmt.Printf("\n")
	}
}

// Try to extend a knight's tour starting at (start_row, start_col).
// Return true or false to indicate whether we have found a solution.
// board stores cells which correspond the index (starting from 0) of the move when the knight was on this square,
// or unvisited (-1) if the knight has not yet visited this square
func find_tour(board [][]int, num_rows, num_cols, cur_row, cur_col, num_visited int) bool {
	num_calls += 1

	if num_visited == num_rows*num_cols {
		// the knight has previously visited every square
		if require_closed_tour {
			// we need to see if this is a closed tour
			for _, offset := range move_offsets {
				new_row := cur_row + offset.dr
				new_col := cur_col + offset.dc
				// if the resulting position is a legal move and ends up in the cell we visited first, then this is a valid closed tour
				if 0 <= new_row && new_row < num_rows && 0 <= new_col && new_col < num_cols && board[new_row][new_col] == 0 {
					return true
				}
			}
			// no possible move produced a closed tour, so because we have visited every square, we are out of options
			return false
		} else {
			// we have been asked to find an open tour, and we have found it
			return true
		}
	}

	// the knight has not visited every square
	for _, offset := range move_offsets {
		new_row := cur_row + offset.dr
		new_col := cur_col + offset.dc
		// the test move is off the board
		if new_row < 0 || new_col < 0 || new_row >= num_rows || new_col >= num_cols {
			continue
		}

		// given that the test move is not off the board, it moves to a previously visited square
		if board[new_row][new_col] != unvisited {
			continue
		}

		// now we know that the test move is viable
		// we are visiting it in the num_visited turn
		board[new_row][new_col] = num_visited

		// see if we can find a tour from this new board arrangement
		tour_result := find_tour(board, num_rows, num_cols, new_row, new_col, num_visited+1)

		// we found a complete tour
		if tour_result {
			return true
		} else {
			// there is no tour starting with the new board arrangement including the test move
			// we must undo the test move
			board[new_row][new_col] = unvisited
		}
	}

	// if we have reached this point, then none of the test moves gives a complete tour,
	// so there is no complete tour starting from the current board position
	return false
}

func main() {
	num_calls = 0

	// Initialize the move offsets.
	initialize_offsets()

	// Create the blank board.
	board := make_board(num_rows, num_cols)

	// Try to find a tour.
	start := time.Now()
	board[0][0] = 0
	if find_tour(board, num_rows, num_cols, 0, 0, 1) {
		fmt.Println("Success!")
	} else {
		fmt.Println("Could not find a tour.")
	}
	elapsed := time.Since(start)
	dump_board(board)
	fmt.Printf("%f seconds\n", elapsed.Seconds())
	fmt.Printf("%d calls\n", num_calls)
}
