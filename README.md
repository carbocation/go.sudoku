go.sudoku
=========

A go implementation of Norvig's Sudoku solver

This will solve any sudoku puzzle that has a valid, single solution. It implements constraint propagation and 
depth-first search. Many "easy" puzzles are solved just by propagating the constraints that come from the 
statement of the puzzle itself. The rest are solved by applying depth-first search, starting from squares with 
the fewest remaining options, digging deeper from there, and backtracking if a contradiction is found.

Patches would be welcome to take advantage of goroutines / to parallelize some of the depth-first search 
and backtracking.
