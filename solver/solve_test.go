package solver

import (
	"testing"
)

// Using several "hard" inputs
var hardMap = map[string]string{
	`....7..2.8.......6.1.2.5...9.54....8.........3....85.1...3.2.8.4.......9.7..6....`: `5 9 4 |8 7 6 |1 2 3 
8 2 3 |9 1 4 |7 5 6 
6 1 7 |2 3 5 |8 9 4 
------+------+------
9 6 5 |4 2 1 |3 7 8 
7 8 1 |6 5 3 |9 4 2 
3 4 2 |7 9 8 |5 6 1 
------+------+------
1 5 9 |3 4 2 |6 8 7 
4 3 6 |5 8 7 |2 1 9 
2 7 8 |1 6 9 |4 3 5 `,
	`7.....4...2..7..8...3..8.799..5..3...6..2..9...1.97..6...3..9...3..4..6...9..1.35`: `7 9 8 |6 3 5 |4 2 1 
1 2 6 |9 7 4 |5 8 3 
4 5 3 |2 1 8 |6 7 9 
------+------+------
9 7 2 |5 8 6 |3 1 4 
5 6 4 |1 2 3 |8 9 7 
3 8 1 |4 9 7 |2 5 6 
------+------+------
6 1 7 |3 5 2 |9 4 8 
8 3 5 |7 4 9 |1 6 2 
2 4 9 |8 6 1 |7 3 5`,
	`.6.5.4.3.1...9...8.........9...5...6.4.6.2.7.7...4...5.........4...8...1.5.2.3.4.`: `8 6 9 |5 7 4 |1 3 2 
1 2 4 |3 9 6 |7 5 8 
3 7 5 |1 2 8 |6 9 4 
------+------+------
9 3 2 |8 5 7 |4 1 6 
5 4 1 |6 3 2 |8 7 9 
7 8 6 |9 4 1 |3 2 5 
------+------+------
2 1 7 |4 6 9 |5 8 3 
4 9 3 |7 8 5 |2 6 1 
6 5 8 |2 1 3 |9 4 7 `,
	`...92......68.3...19..7...623..4.1....1...7....8.3..297...8..91...5.72......64...`: `3 8 7 |9 2 6 |4 1 5 
5 4 6 |8 1 3 |9 7 2 
1 9 2 |4 7 5 |8 3 6 
------+------+------
2 3 5 |7 4 9 |1 6 8 
9 6 1 |2 5 8 |7 4 3 
4 7 8 |6 3 1 |5 2 9 
------+------+------
7 5 4 |3 8 2 |6 9 1 
6 1 3 |5 9 7 |2 8 4 
8 2 9 |1 6 4 |3 5 7 
`,
	`1...34.8....8..5....4.6..21.18......3..1.2..6......81.52..7.9....6..9....9.64...2`: `1 5 2 |9 3 4 |6 8 7 
7 6 3 |8 2 1 |5 4 9 
9 8 4 |5 6 7 |3 2 1 
------+------+------
6 1 8 |4 9 3 |2 7 5 
3 7 5 |1 8 2 |4 9 6 
2 4 9 |7 5 6 |8 1 3 
------+------+------
5 2 1 |3 7 8 |9 6 4 
4 3 6 |2 1 9 |7 5 8 
8 9 7 |6 4 5 |1 3 2 `,
	`...8.1..........435............7.8........1...2..3....6......75..34........2..6..`: `2 3 7 |8 4 1 |5 6 9 
1 8 6 |7 9 5 |2 4 3
5 9 4 |3 2 6 |7 1 8 
------+------+------
3 1 5 |6 7 4 |8 9 2 
4 6 9 |5 8 2 |1 3 7 
7 2 8 |1 3 9 |4 5 6 
------+------+------
6 4 2 |9 1 8 |3 7 5 
8 5 3 |4 6 7 |9 2 1 
9 7 1 |2 5 3 |6 8 4 `,
}

func TestSolve(t *testing.T) {

	for u, s := range hardMap {
		unsolved, err := Solve(u)
		if err != nil {
			t.Errorf(`Solving the unsolved grid yielded an error (%v):%v\n`, err, u)
			continue
		}
		solved, err := ParseGrid(s)
		if err != nil {
			t.Errorf(`The "solved" grid was illegal (%v) and could not be parsed:%v\n`, err, s)
			continue
		}

		if x, y := Display(unsolved), Display(solved); x != y {
			t.Errorf("The following two should be equal:%v\n%v", x, y)
			continue
		}
	}
}
