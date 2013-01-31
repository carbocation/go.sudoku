package main

import (
    "github.com/carbocation/sudoku.git/sudoku"
    "fmt"
)

func main() {
    digits := "123456789"
    rows := "ABCDEFGHI"
    squares := sudoku.Cross(rows, digits)
    
    
    fmt.Println(squares)
}