/*
Package solver implements the sudoku solver written about in Peter Norvig's essay,
"Solving Every Sudoku Puzzle" http://norvig.com/sudoku.html

Go implementation by James Pirruccello, 2013
*/
package solver

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

//The name of a square, which is the discrete entity that needs to be filled
type Square string

//The possible values that remain for each Square
type SquarePossibilities map[Square]string

var digits, rows, cols = "123456789", "ABCDEFGHI", "123456789"
var squares []Square
var unitlist [][]Square
var peers map[Square][]Square
var units map[Square][][]Square
var validCharRegex *regexp.Regexp

func init() {
	initialize()
}

func initialize() {
	validCharRegex = regexp.MustCompile(`[0-9]*\.*`)
	squares = Cross(rows, cols)
	unitlist = BuildUnitList(rows, cols, []string{"ABC", "DEF", "GHI"}, []string{"123", "456", "789"})
	units = map[Square][][]Square{}
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

	peers = map[Square][]Square{}
	for square, unitlist := range units {
		for _, unit := range unitlist {
		NextUnitSquare:
			for _, unitsquare := range unit {
				if unitsquare != square {
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
}

//Take two strings and build a slice of strings, each of which are
//bigrams taken from every AxB combination
func Cross(A string, B string) []Square {
	var out []Square
	for _, i := range A {
		for _, j := range B {
			out = append(out, Square(i)+Square(j))
		}
	}

	return out
}

func BuildUnitList(rows string, cols string, rowBlocks []string, rowCols []string) [][]Square {
	out := [][]Square{}
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
func ParseGrid(Grid string) (SquarePossibilities, error) {
	//At initialization, `values` says "every square can contain every value"
	values := SquarePossibilities{}
	for _, square := range squares {
		values[square] = digits
	}

	//If you can't parse the grid, the game makes no sense
	vGrid := strings.Join(validCharRegex.FindAllString(Grid, -1), "")
	gridMap, err := gridSquarePossibilities(vGrid)
	if err != nil {
		return nil, err
	}

	//For each square and its value(s) from the GridMap
	for s, d := range gridMap {
		//Test each allowed digit
		for _, xd := range digits {
			//If the GridMap square's value is an allowed digit (not just a blank placeholder)
			if d == string(xd) {
				//Try to assign that value
				values, err = assign(values, s, d)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return values, nil
}

//Convert grid into a dict of {square: char} with '0' or '.' for empties.
func gridSquarePossibilities(vGrid string) (SquarePossibilities, error) {
	out := SquarePossibilities{}

	//Accept any grid representation, but now let's slim the grid down to just the valid characters.
	if len(vGrid) != 81 {
		return SquarePossibilities{}, errors.New("vGrid does not contain 81 valid digits.")
	}

	//Now create our output map, where each square on the grid is filled in with an empty placeholder or a value
	for i := 0; i < 81; i++ {
		if string(vGrid[i]) == "0" {
			out[squares[i]] = "."
		} else {
			out[squares[i]] = string(vGrid[i])
		}
	}

	return out, nil
}

//Assign square s to digit d, and propagate.
func assign(values SquarePossibilities, s Square, d string) (SquarePossibilities, error) {
	//If this square has no possible values, fail
	if len(values[s]) < 1 {
		return nil, errors.New(fmt.Sprintf("values[%s] has no permissible digits left; contradiction.", s))
	}

	//Find all other digit values that this square previously could have accepted.
	otherSquarePossibilities := ``
	for _, v := range values[s] {
		//This value, d, is by definition not one of the "other" values
		if string(v) == d {
			continue
		}

		otherSquarePossibilities = otherSquarePossibilities + string(v)
	}

	//Now try to eliminate all of these other values from the master values record
	if len(otherSquarePossibilities) > 0 {
		for _, d2 := range otherSquarePossibilities {
			if _, err := eliminate(values, s, string(d2)); err != nil {
				return nil, err
			}
		}
	}

	//Default: succeed
	return values, nil
}

//Eliminate digit d from the list of possible values at square s (values[s]) and propagate.
func eliminate(values SquarePossibilities, s Square, d string) (SquarePossibilities, error) {
	err := error(nil)

	dInSquarePossibilitiesS := false
	for _, val := range values[s] {
		if string(val) == d {
			dInSquarePossibilitiesS = true
			break
		}
	}

	//Digit has already been eliminated from values[s]. We're done here.
	if !dInSquarePossibilitiesS {
		return values, nil
	}

	//Remove digit from values[s]
	values[s] = strings.Replace(values[s], d, "", -1)

	//Having removed the digit from the square's possible values, if values[s] is now empty,
	// we've arrived at a contradiction.
	if len(values[s]) == 0 {
		//Contradiction: removed last possible value for this square
		return nil, errors.New(fmt.Sprintf("Cannot eliminate %s from values[%s] because now values[%s] has no valid potential digits (valid: %s).", d, s, s, values[s]))
	} else if len(values[s]) == 1 {
		//(1) If this square s is reduced to one possible digit d2, then that is the square's solution. Eliminate digit d2 from its peers.
		d2 := values[s]
		for _, s2 := range peers[s] {
			//Eliminate the only possible answer to values[s] from all of the peers of s
			if _, err = eliminate(values, s2, d2); err != nil {
				return nil, err
			}
		}
	}

	//(2) For all of the units that this square belongs to,
	//if the unit is reduced to only one place where this digit d could go, then put d there.
	for _, u := range units[s] {
		//Iterate over all of the other squares in this square's units.
		// For this unit, grab all of the squares that can accept `d` (enumerate them in dPlaces)
		dPlaces := []Square{}
		//For every square in the unit
		for _, s2 := range u {
			//For every digit that this square can accept
			for _, d2 := range values[s2] {
				//If d is in that square's digit list
				if d == string(d2) {
					//Then dPlaces includes this square
					dPlaces = append(dPlaces, s2)
					break
				}
			}
		}

		if len(dPlaces) == 0 {
			//No place to put d; contradiction
			return values, errors.New(fmt.Sprintf("There is no place in unit %s to put %s.", u, d))
		} else if len(dPlaces) == 1 {
			//D must go into dPlaces[0]
			_, err = assign(values, dPlaces[0], d)
			if err != nil {
				//If it can't, you'll have to backtrack (or you have an insoluble puzzle, depending on nesting depth)
				return nil, err
			}
		}
	}

	return values, nil
}

func Solve(Grid string) (SquarePossibilities, error) {
	values, err := ParseGrid(Grid)
	if err != nil {
		return values, err
	}

	return search(values, nil)
}

func search(values SquarePossibilities, err error) (SquarePossibilities, error) {
	//Exit: already failed
	if err != nil {
		return values, err
	}

	//Exit: all squares have exactly one possible value, so you have already solved the puzzle
	solved := true
	for _, s := range squares {
		if len(values[s]) != 1 {
			solved = false
		}
	}
	if solved {
		return values, nil
	}

	//The puzzle is not yet solved. For each square with more than 1 possibility,
	// try to assign each possible digit to it. Then proceed depth-first unless you
	// encounter a contradiction. If you do, back off and try the next square.

	//Chose the unfilled square s with the fewest possibilities
	//(minimize probability of guessing wrong)
	min, sq := len(digits)+1, Square("")
	for _, s := range squares {
		if len(values[s]) < min && len(values[s]) > 1 {
			min = len(values[s])
			sq = s

			if min == 2 {
				break
			}
		}
	}

	//For every possible digit, try to assign it to this square
	//This recursive call to search is the heart of the depth-first search
	for _, d := range values[sq] {
		//*Cloning `values` is important*; if you don't, you mutate the original map and
		// then you cannot backtrack if your depth-search algorithm fails to
		// guess right the first time at every step
		//ch make(<-map[string]string
		//go func() (map[string]string, error) {
		vCloned, err := assign(cloneSquarePossibilities(values), sq, string(d))
		//fmt.Println("Test: Assigning sq", sq, "with options", values[sq], "to d", string(d), "yielded", err)
		if err != nil {
			continue
		}
		vCloned, err = search(vCloned, err)
		if err == nil {
			return vCloned, nil
		}
		//}()
	}

	return nil, errors.New("Your depth-first search failed on this branch.")
}

//Clone the values map
func cloneSquarePossibilities(values SquarePossibilities) SquarePossibilities {
	cpySquarePossibilities := make(SquarePossibilities, len(values))
	for k, v := range values {
		cpySquarePossibilities[k] = v
	}

	return cpySquarePossibilities
}

func DisplayCompact(values map[string]string) string {
	out := []byte{}
	for _, r := range rows {
		for _, c := range cols {
			s := fmt.Sprintf("%s%s", string(r), string(c))
			out = append(out, []byte(values[s])...)
		}
	}

	return string(out)
}

func Display(values SquarePossibilities) string {
	mLen := 0
	for _, s := range squares {
		if len(values[s]) > mLen {
			mLen = len(values[s])
		}
	}
	width := mLen + 1

	line := []byte{}
	for loop := 0; loop < 3; loop++ {
		for i := 0; i < width*3; i++ {
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
			s := Square(fmt.Sprintf("%s%s", string(r), string(c)))
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
