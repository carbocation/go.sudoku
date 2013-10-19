/*
Package sudoku implements the sudoku solver written about in Peter Norvig's essay,
"Solving Every Sudoku Puzzle" http://norvig.com/sudoku.html

Go implementation by James Pirruccello, 2013
*/
package sudoku

import (
	//"bytes"
	"regexp"
	"strings"
)

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

//Convert any representation of a grid into a valid one for our internal representation
func ParseGrid(Grid string) string {
	re := regexp.MustCompile(`[0-9]*\.*`)
	return strings.Join(re.FindAllString(Grid, -1), "")
}
