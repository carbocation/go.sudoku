package main

import (
    "github.com/carbocation/sudoku.git/sudoku"
    "fmt"
)

func main() {
	if pg, ok := sudoku.ParseGrid("9...............9..............4...........5.............1............3.........2"); ok {
		fmt.Println("Parsed Grid: ", pg)
	} else {
		fmt.Println("Parsed Grid: Illegal.", len(pg), " valid chars (must be 81 valid chars)")
	}
	
	/*
	if pg, ok := sudoku.ParseGrid("9..............|||................................................................2"); ok {
		fmt.Println("Parsed Grid: ", pg)
	} else {
		fmt.Println("Parsed Grid: Illegal.", len(pg), " valid chars (must be 81 valid chars)")
	}
	*/
}
