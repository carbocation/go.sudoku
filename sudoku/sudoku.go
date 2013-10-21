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
//var values map[string]string
var peers map[string][]string
var units map[string][][]string
var validCharRegex *regexp.Regexp

func init() {
	initialize()
}

func initialize() {
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
func ParseGrid(Grid string) (map[string]string, bool) {
	//initialize()
	
	//At initialization, `values` says "every square can contain every value"
	values := map[string]string{}
	for _, square := range squares {
		values[square] = digits
	}
	
	//If you can't parse the grid, the game makes no sense
	vGrid := strings.Join(validCharRegex.FindAllString(Grid, -1), "")
	gridMap, ok := gridValues(vGrid)
	if !ok {
		return values, false
	}
	
	//For each square and its value from the GridMap 
	for s, d := range gridMap {
		//For each allowed digit
		for _, xd := range digits {
			//If the GridMap square's value is an allowed digit
			if d == string(xd) && !assign(values, s, d) {
				return values, false
			}
		}
	}
	
	return values, true
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

//Eliminate all the other possible values at square s, except digit d, and propagate. 
// Return false if a contradiction is detected.
func assign(values map[string]string, s, d string) bool {
	//Find all other digit values that this square previously could have accepted.
	otherValues := make([]string, len(values[s])-1)
	for _, v := range values[s] {
		if string(v) == d {
			continue
		}
		
		otherValues = append(otherValues, string(v))
	}
	
	//Now eliminate all of these values from the master values record
	//If the operation ever fails, return with failure
	for _, d2 := range otherValues {
		if ok := eliminate(values, s, d2); !ok {
			return false
		}
	}
	
	//Default: succeed
	return true
}

//Eliminate digit d from the list of possible values at square s (values[s]); propagate when values or places <= 2.
//  Return false if a contradiction is detected.
func eliminate(values map[string]string, s, d string) bool {
	
	dInValuesS, dKey := false, 0
	for dk, val := range values[s] {
		if string(val) == d {
			dInValuesS = true
			dKey = dk
			break
		}
	}
	
	//Digit has already been eliminated from values[s]. We're done here.
	if !dInValuesS {
		return true
	}
	
	//Remove digit from values[s]
	values[s] = values[s][:dKey] + values[s][dKey+1:]
	
	//Having removed the digit from the square's possible values, if values[s] is now empty, 
	// we've arrived at a contradiction.
	if len(values[s]) == 0 {
		//Contradiction: removed last possible value for this square
		//fmt.Println("Cannot remove ",d," from values[s] because that leaves you with 0 possible values.") 
		return false
	} else if len(values[s]) == 1 {
		//(1) If this square s is reduced to one possible digit d2, then that is the square's solution. Eliminate digit d2 from its peers.
		d2 := values[s] //The only remaining value `d2` for values[s]
		//fmt.Println("Assigned values[",s,"]==", d2,", its final value. Now eliminating ", d2," from peers[", s,"] == ", peers[s]) 
		for _, s2 := range peers[s] {
			//Eliminate the only possible answer to values[s] from all of the peers of s
			if ok := eliminate(values, s2, d2); !ok {
				return false
			}
		}
	}
	
	//(2) For all of the units that this square belongs to, 
	//if the unit is reduced to only one place where this digit d could go, then put d there.
	for _, u := range units[s] {
		//fmt.Println("units[", s,"] == ", u)
		/*
		fmt.Println("values[s] = ", values[s])
		fmt.Println(d)
		*/
		//Iterate over all of the other squares in this square's units. 
		// For this unit, grab all of the squares that can accept `d`
		dPlaces := []string{}
		//For every square in the unit
		for _, s2 := range u {
			NextSquare:
			//For every digit that this square can accept
			for _, d2 := range values[s2] {
				//If d is in that square's digit list
				if d == string(d2) {
					//Then dPlaces includes this square
					dPlaces = append(dPlaces, s2)
					continue NextSquare 
				}
			}
		}
		
		if len(dPlaces) == 0 {
			//No place to put d; contradiction
			//fmt.Println("Fail: no place to put digit ", d," in unit ", u)
			return false
		} else if len(dPlaces) == 1 {
			//D must go into dPlaces[0]
			//fmt.Println("While ",d," was being excluded from ", s, ", it was discovered that for group ", u," ", d, " could only belong to ", dPlaces[0], ". ")
			if !assign(values, dPlaces[0], d) {
				//fmt.Println("Some problem occurred when assigning d to dPlaces[0]")
				return false 
			}
		}
	}
	
	return true
}

//func Solve

func DisplayCompact(values map[string]string) string {
	out := []byte{}
	for _, r := range rows {
		for _, c := range cols {
			s := fmt.Sprintf("%s%s", string(r),string(c))
			out = append(out, []byte(values[s])...)
		}
	}
	
	return string(out)
}

func Display(values map[string]string) string {
	mLen := 0
	for _, s := range squares { 
		if len(values[s]) > mLen {
			mLen = len(values[s])
		}
	}
	width := mLen + 1
	
	line := []byte{}
	for loop := 0; loop < 3; loop++ {
		for i := 0; i < width * 3; i++ {
			line = append(line, []byte("-")...)
		}
		if loop < 2 {
			line = append(line, []byte("+")...)
		}
	}
	line = append(line, []byte("\n")...)
	
	
	out := []byte{}
	for rn, r := range rows {
		for cn, c := range cols {
			s := fmt.Sprintf("%s%s", string(r),string(c))
			out = append(out, []byte(strWid(values[s], width))...)
			
			if cn == 2 || cn == 5 {
				out = append(out, []byte(`|`)...)
			}
		}
		out = append(out, []byte("\n")...)
		if rn == 2 || rn == 5 {
			out = append(out, line...)
		}
	}
	
	return string(out)
}

func strWid(s string, width int) string {
	if len(s) >= width {
		return s
	}
	
	for i := len(s); i < width; i++ {
		s += " "
	}
	
	return s 
}
