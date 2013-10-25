package main

import (
	"fmt"
	"strings"

	"github.com/carbocation/go.sudoku/sudoku"
)

var values, result, err = sudoku.SquarePossibilities{}, sudoku.SquarePossibilities{}, error(nil)
//Pathologic: .....6....59.....82....8....45........3........6..3.54...325..6..................
var hardest = strings.Split(`...8.1..........435............7.8........1...2..3....6......75..34........2..6..
85...24..72......9..4.........1.7..23.5...9...4...........8..7..17..........36.4.
..53.....8......2..7..1.5..4....53...1..7...6..32...8..6.5....9..4....3......97..
12..4......5.69.1...9...5.........7.7...52.9..3......2.9.6...5.4..9..8.1..3...9.4
...57..3.1......2.7...234......8...4..7..4...49....6.5.42...3.....7..9....18.....
7..1523........92....3.....1....47.8.......6............9...5.6.4.9.7...8....6.1.
1....7.9..3..2...8..96..5....53..9...1..8...26....4...3......1..4......7..7...3..
1...34.8....8..5....4.6..21.18......3..1.2..6......81.52..7.9....6..9....9.64...2
...92......68.3...19..7...623..4.1....1...7....8.3..297...8..91...5.72......64...
.6.5.4.3.1...9...8.........9...5...6.4.6.2.7.7...4...5.........4...8...1.5.2.3.4.
7.....4...2..7..8...3..8.799..5..3...6..2..9...1.97..6...3..9...3..4..6...9..1.35
....7..2.8.......6.1.2.5...9.54....8.........3....85.1...3.2.8.4.......9.7..6....`, "\n")

func main() {
	for _, input := range hardest {
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		fmt.Println("Input: ")
		fmt.Println(input)
		if values, err = sudoku.ParseGrid(input); err == nil {
			fmt.Println("Grid is valid.")
		} else {
			fmt.Println("Parsed Grid: Illegal.", len(values), " valid chars (must be 81 valid chars)")
			fmt.Println(err)
		}
		fmt.Println("Loaded: ")
		fmt.Println(sudoku.Display(values))
		fmt.Println("Solved:")
		result, err = sudoku.Solve(input)
		if err != nil {
			fmt.Println("Some unknown error during solving.")
			fmt.Println(err)
		} else {
			fmt.Println(sudoku.Display(result))
		}
		fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
	}
}
