package main

import (
    "github.com/carbocation/sudoku.git/sudoku"
    "fmt"
)

var values, ok, input = map[string]string{}, false, ``

func main() {
	input = 
`8 5 . |. . 2 |4 . . 
7 2 . |. . . |. . 9 
. . 4 |. . . |. . . 
------+------+------
. . . |1 . 7 |. . 2 
3 . 5 |. . . |9 . . 
. 4 . |. . . |. . . 
------+------+------
. . . |. 8 . |. 7 . 
. 1 7 |. . . |. . . 
. . . |. 3 6 |. 4 . `
	fmt.Println("Input: ")
	fmt.Println(input)
	if values, ok = sudoku.ParseGrid(input); ok {
		fmt.Println("Grid is valid.")
	} else {
		fmt.Println("Parsed Grid: Illegal.", len(values), " valid chars (must be 81 valid chars)")
	}
	fmt.Println("Compact: ", sudoku.DisplayCompact(values))
	fmt.Println("Full:\n")
	result, ok := sudoku.Solve(input)
	if !ok {
		fmt.Println("Some unknown error during solving.")
	}
	fmt.Println(sudoku.Display(result))
	
	
/*
	input = 
`0 0 3 |0 2 0 |6 0 0 
9 0 0 |3 0 5 |0 0 1 
0 0 1 |8 0 6 |4 0 0 
------+------+------
0 0 8 |1 0 2 |9 0 0 
7 0 0 |0 0 0 |0 0 8 
0 0 6 |7 0 8 |2 0 0 
------+------+------
0 0 2 |6 0 9 |5 0 0 
8 0 0 |2 0 3 |0 0 9 
0 0 5 |0 1 0 |3 0 0  `
	fmt.Println("Input: ")
	fmt.Println(input)
	if values, ok = sudoku.ParseGrid(input); ok {
		fmt.Println("Grid is valid.")
	} else {
		fmt.Println("Parsed Grid: Illegal.", len(values), " valid chars (must be 81 valid chars)")
	}
	fmt.Println("Compact: ", sudoku.DisplayCompact(values))
	fmt.Println("Full:\n")
	fmt.Println(sudoku.Display(values))
*/
}
