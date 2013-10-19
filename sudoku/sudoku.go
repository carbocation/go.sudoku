package sudoku

import (
    //"bytes"
)

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