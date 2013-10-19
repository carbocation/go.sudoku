package main

import (
    "github.com/carbocation/sudoku.git/sudoku"
    "fmt"
)

func main() {
    digits := "123456789"
    rows := "ABCDEFGHI"
    cols := digits
    
    squares := sudoku.Cross(rows, cols)
    fmt.Println("squares: ", squares)
    unitlist := sudoku.BuildUnitList(rows, cols, []string{"ABC", "DEF", "GHI"}, []string{"123", "456", "789"})
	fmt.Println("unitlist: ", unitlist)
	units := map[string][][]string{}
	for _, square := range squares {
		Square:
		for _, v := range unitlist {
			for _, unitsquare := range v {
				//If the square is contained within the unitlist, then add this unitlist belongs to this square's unit map 
				if square == unitsquare {
					units[square] = append(units[square], v)
					continue Square
				}
			}
		}
	}
	fmt.Println("units: ", units)
	
	peers := map[string][]string{}
	for square, unitlist := range units {
		for _, unit := range unitlist {
			NextUnitSquare:
			for _, unitsquare := range unit {
				if unitsquare != square {
					//TODO: if the unitsquare is already in peers[square], don't re-add it.
					for _, ps := range peers[square] {
						if ps == unitsquare {
							continue NextUnitSquare 
						}
					}
					
					peers[square] = append(peers[square], unitsquare)
				}
			}
		}
	}
	fmt.Println("peers[C2]: ", peers["C2"])
	
	panic(0)
	fmt.Println(sudoku.ParseGrid("324|.342|ajfdso28.49....2"))
}