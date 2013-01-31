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