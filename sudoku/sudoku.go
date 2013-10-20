/*
Package sudoku implements the sudoku solver written about in Peter Norvig's essay,
"Solving Every Sudoku Puzzle" http://norvig.com/sudoku.html

Go implementation by James Pirruccello, 2013
*/
package sudoku

import (
	//"bytes"
	"regexp"
	"fmt"
	"strings"
)

var digits, rows, cols = "123456789", "ABCDEFGHI", "123456789"
var squares []string
var unitlist [][]string
var peers map[string][]string
var units map[string][][]string
var validCharRegex *regexp.Regexp

func init() {
	validCharRegex = regexp.MustCompile(`[0-9]*\.*`)
    squares = Cross(rows, cols)
    unitlist = BuildUnitList(rows, cols, []string{"ABC", "DEF", "GHI"}, []string{"123", "456", "789"})
	units = map[string][][]string{}
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
	
	peers = map[string][]string{}
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
	
	/*
	fmt.Println("squares: ", squares)
	fmt.Println("unitlist: ", unitlist)
	fmt.Println("units: ", units)
	fmt.Println("peers[C2]: ", peers["C2"])
	*/
}

//Take two strings and build a slice of strings, each of which are
//bigrams taken from every AxB combination
func Cross(A string, B string) []string {
	var out []string
	for _, i := range A {
		for _, j := range B {
			out = append(out, string(i)+string(j))
		}
	}

	return out
}

func BuildUnitList(rows string, cols string, rowBlocks []string, rowCols []string) [][]string {
	out := [][]string{}
	for _, c := range cols {
		out = append(out, Cross(rows, string(c)))
	}

	for _, r := range rows {
		out = append(out, Cross(string(r), cols))
	}

	for _, rs := range rowBlocks {
		for _, cs := range rowCols {
			out = append(out, Cross(rs, cs))
		}
	}

	return out
}

//Convert grid to a map of possible values, {square: digits}, or emit false if a contradiction is detected.
func ParseGrid(Grid string) (string, bool) {
	//To start, every square can be any digit; then assign values from the grid.
	//At initialization, `values` says "every square can contain every value"
	values := map[string][]string{}
	for _, square := range squares {
		values[square] = append(values[square], digits)
	}
	
	//If you can't parse the grid, the game makes no sense
	vGrid := strings.Join(validCharRegex.FindAllString(Grid, -1), "")
	if gridMap, ok := gridValues(vGrid); !ok {
		return vGrid, false
	} else {
		return fmt.Sprintln(gridMap), true
		
		//TODO: 
		for s, d := range gridMap {
			//etc
		} 
	}
}

//Convert grid into a dict of {square: char} with '0' or '.' for empties.
func gridValues(vGrid string) (map[string]string, bool) {
	out := map[string]string{}
	
	//Accept any grid representation, but now let's slim the grid down to just the valid characters.
	if len(vGrid) != 81 {
		return map[string]string{}, false
	}
	
	//Now create our output map, where each square on the grid is filled in with an empty placeholder or a value
	for i := 0; i < 81; i++ {
		out[squares[i]] = string(vGrid[i])
	}
	
	return out, true
}
