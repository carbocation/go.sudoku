package main

import (
    "github.com/carbocation/sudoku.git/sudoku"
    "fmt"
)

func main() {
	input := ``
	
	input = "912.587.6.3.....9...8..9..5.7..4....6......5....9....7...1..4...5.....3...4..6.82"
	fmt.Println("Input: ")
	fmt.Println(input)
	if pg, ok := sudoku.ParseGrid(input); ok {
		fmt.Println("Grid is valid. Parsed grid: ", pg)
	} else {
		fmt.Println("Parsed Grid: Illegal.", len(pg), " valid chars (must be 81 valid chars)")
	}
	fmt.Println("Compact: ", sudoku.DisplayCompact())
	fmt.Println("Full:\n")
	fmt.Println(sudoku.Display())
	
	
	
	input = `0 0 3 |0 2 0 |6 0 0 
9 0 0 |3 0 5 |0 0 1 
0 0 1 |8 0 6 |4 0 0 
------+------+------
0 0 8 |1 0 2 |9 0 0 
7 0 0 |0 0 0 |0 0 8 
0 0 6 |7 0 8 |2 0 0 
------+------+------
0 0 2 |6 0 9 |5 0 0 
8 0 0 |2 0 3 |0 0 9 
0 0 5 |0 1 0 |3 0 0 `
	//input = "003020600900305001001806400008102900700000008006708200002609500800203009005010300"
	fmt.Println("Input: ")
	fmt.Println(input)
	if pg, ok := sudoku.ParseGrid(input); ok {
		fmt.Println("Grid is valid. Parsed grid: ", pg)
	} else {
		fmt.Println("Parsed Grid: Illegal.", len(pg), " valid chars (must be 81 valid chars)")
	}
	
	
	fmt.Println("Compact: ", sudoku.DisplayCompact())
	fmt.Println("Full:\n")
	fmt.Println(sudoku.Display())
	
}
