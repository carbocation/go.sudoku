package main

import (
    "github.com/carbocation/sudoku.git/sudoku"
    "fmt"
)

func main() {
	
	
	if pg, ok := sudoku.ParseGrid("912.587.6.3.....9...8..9..5.7..4....6......5....9....7...1..4...5.....3...4..6.82"); ok {
		fmt.Println("Grid is valid. Parsed grid: ", pg)
	} else {
		fmt.Println("Parsed Grid: Illegal.", len(pg), " valid chars (must be 81 valid chars)")
	}
	
	
	
	if pg, ok := sudoku.ParseGrid("003020600900305001001806400008102900700000008006708200002609500800203009005010300"); ok {
		fmt.Println("Grid is valid. Parsed grid: ", pg)
	} else {
		fmt.Println("Parsed Grid: Illegal.", len(pg), " valid chars (must be 81 valid chars)")
	}
	
}
