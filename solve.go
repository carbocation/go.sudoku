package main

import (
    "github.com/carbocation/sudoku.git/sudoku"
    "fmt"
)

func main() {
    digits := "123456789"
    rows := "ABCDEFGHI"
    cols := digits
    //squares := sudoku.Cross(rows, cols)
    unitlist := sudoku.BuildUnitList(rows, cols, []string{"ABC", "DEF", "GHI"}, []string{"123", "456", "789"})

    

    fmt.Println(unitlist)
}