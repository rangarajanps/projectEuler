package main

import (
	"fmt"
	"math"
	"strings"
)

type solutionSpace struct {
	row       int
	col       int
	count     int
	possibles []int
}

type sudoku struct {
	gridSize      int
	rows          [][]int
	cols          [][]int
	cubes         [][]int
	filledRow     []map[int]bool
	filledCol     []map[int]bool
	filledCube    []map[int]bool
	filledNumbers map[int]int
}

func NewSudoku(inp string) *sudoku {
	s := new(sudoku)
	s.filledRow = make([]map[int]bool, 9)
	s.filledCol = make([]map[int]bool, 9)
	return s
}

func New(inp string) *sudoku {

	m := NewSudoku(inp)

	gridLength := int(math.Sqrt(float64(len(inp))))
	m.gridSize = gridLength

	// initialize 2 grids - one for row wise access and other for column wise
	g := make([][]int, gridLength)
	for i := range g {
		g[i] = make([]int, gridLength)
	}
	t := make([][]int, gridLength)
	for i := range t {
		t[i] = make([]int, gridLength)
	}
	// c := make([][]int, gridLength)
	// for i := range c {
	// 	c[i] = make([]int, gridLength)
	// }
	// m.cubes = c

	rowsToFill := make([]map[int]bool, m.gridSize)
	colsToFill := make([]map[int]bool, m.gridSize)
	cubesToFill := make([]map[int]bool, m.gridSize)
	for i := 0; i < m.gridSize; i++ {
		nmr := make(map[int]bool)
		nmc := make(map[int]bool)
		nmcube := make(map[int]bool)
		for n := 1; n < 10; n++ {
			nmr[n] = false
			nmc[n] = false
			nmcube[n] = false
		}
		rowsToFill[i] = nmr
		colsToFill[i] = nmc
		cubesToFill[i] = nmcube
	}
	m.filledRow = rowsToFill
	m.filledCol = colsToFill
	m.filledCube = cubesToFill

	numbersToFill := make(map[int]int)

	rowSl := make([]string, gridLength)
	arr := []rune(inp)

	// rowToUpdate := m.filledRow
	// colToUpdate := m.filledCol

	for row := 0; row < gridLength; row++ {
		startIndex := row * gridLength
		rowSl[row] = string(arr[startIndex : startIndex+9])
		colSl := []rune(rowSl[row])
		for col := 0; col < gridLength; col++ {
			value := int(colSl[col]) - 48
			numbersToFill[value]++
			g[row][col] = value
			t[col][row] = value
			if value != 0 {
				// fmt.Println("Updating the map for ", value, m.filledRow[0])
				m.filledRow[row][value] = true
				m.filledCol[col][value] = true
			}

		}
	}

	// rows := make([][]int, gridLength)
	// for row := 0; row < gridLength; row++ {
	// 	rows[row] = g[row]
	// }

	cols := make([][]int, gridLength)
	for row := 0; row < gridLength; row++ {
		cols[row] = t[row]
	}

	m.rows = g
	m.cols = cols
	m.cubes = m.initializeCubeWiseList(g)
	fmt.Println("Numbers filled ", numbersToFill)

	return m
}

func (p sudoku) printNumbersToFill() {

	fmt.Println("Row wise map details ")
	for _, val := range p.filledRow {
		fmt.Println(val)
	}

	fmt.Println("Column wise map details ")
	for _, val := range p.filledCol {
		fmt.Println(val)
	}

	fmt.Println("Cube wise map details ")
	for _, val := range p.filledCube {
		fmt.Println(val)
	}
}

func (p sudoku) solve() {

	// do union to find numbers to fill in that row, col pillar
	// eliminate based on presence of those numbers in other rows, col, grids

}

func (p sudoku) generatePossibleSolutions() {

	grid := p.rows

	var possibleSolutions []solutionSpace

	for row := 0; row < p.gridSize; row++ {
		for col := 0; col < p.gridSize; col++ {
			if grid[row][col] == 0 {
				cubeNum := (row / 3) + (col / 3)
				if row > 2 && row < 6 {
					cubeNum = cubeNum + 2
				}
				if row > 5 {
					cubeNum = cubeNum + 4
				}
				tempSet := union(p.filledRow[row], p.filledCol[col], p.filledCube[cubeNum])
				var numList []int
				for val, isFilled := range tempSet {
					if !isFilled {
						numList = append(numList, val)
					}
				}
				if len(numList) == 0 {
					fmt.Println("Spot to fill first ", row, col, cubeNum, tempSet)
					fmt.Println("Row to check ", p.filledRow[row])
					fmt.Println("Column to check ", p.filledCol[col])
					fmt.Println("Cube to check ", p.filledCube[cubeNum])
					fmt.Println()
				}
				ss := solutionSpace{
					row, col, len(numList), numList,
				}
				possibleSolutions = append(possibleSolutions, ss)
			}
		}
	}
	fmt.Println(possibleSolutions)

}

func (p sudoku) initializeCubeWiseList(grid [][]int) [][]int {

	// grid := p.rows

	cube := make([][]int, 9)
	for i := range cube {
		cube[i] = make([]int, 9)
	}

	for gridRow := 0; gridRow < 3; gridRow++ {

		for gridCol := 0; gridCol < 3; gridCol++ {

			rowIndex := gridRow * 3
			colIndex := (gridCol*3 + gridCol/3) % 9
			gridNumber := rowIndex + colIndex/3
			// fmt.Println("Printing Grid ", gridNumber, rowIndex, colIndex)

			var iList []int
			for r := rowIndex; r < rowIndex+3; r++ {
				for c := colIndex; c < colIndex+3; c++ {
					if grid[r][c] != 0 {
						p.filledCube[gridNumber][grid[r][c]] = true
					}
					iList = append(iList, grid[r][c])
					// fmt.Printf("%d ", grid[r][c])
				}
				// fmt.Println()
			}
			cube[gridNumber] = iList
		}
	}

	// fmt.Println("Cube is ", cube)

	return cube

}

func (p sudoku) calculateCountOfCellsToFill() {

	grid := p.rows

	valuesToFillInRow := make([]int, p.gridSize)
	valuesToFillInCol := make([]int, p.gridSize)
	for row := 0; row < p.gridSize; row++ {
		for col := 0; col < p.gridSize; col++ {
			if grid[row][col] == 0 {
				valuesToFillInRow[row]++
				valuesToFillInCol[col]++
			}
		}
	}
	fmt.Println(valuesToFillInRow)
	fmt.Println(valuesToFillInCol)

}

func main() {
	// test := []int{1, 2, 3, 4, 5}
	// fmt.Println(test)
	inp := `003020600900305001001806400008102900700000008006708200002609500800203009005010300`
	p := New(inp)
	// p.printNumbersToFill()

	prettyPrint(p.rows)
	p.generatePossibleSolutions()
	// prettyPrint(p.cols)
	// prettyPrint(p.cubes)
	// p.solve()
}

func isSolved(grid [][]int) bool {

	gridLength := len(grid)

	sumRow := make([]int, gridLength)
	sumCol := make([]int, gridLength)
	for row := 0; row < gridLength; row++ {
		for col := 0; col < gridLength; col++ {
			sumRow[row] += grid[row][col]
			sumCol[col] += grid[row][col]
		}
	}
	// fmt.Println(sumRow)
	// fmt.Println(sumCol)
	for i := 0; i < gridLength; i++ {
		if sumRow[i] != 45 || sumCol[i] != 45 {
			return false
		}
	}

	return true
}

func union(s1, s2, s3 map[int]bool) map[int]bool {

	s_union := map[int]bool{}
	for i := 1; i < 10; i++ {
		if s1[i] || s2[i] || s3[i] {
			s_union[i] = true
		} else {
			s_union[i] = false
		}
	}

	// fmt.Println(s_union)
	return s_union
}

func genGrid(inp string) [][]int {

	gridLength := int(math.Sqrt(float64(len(inp))))

	g := make([][]int, gridLength)
	for i := range g {
		g[i] = make([]int, gridLength)
	}

	rowSl := make([]string, gridLength)
	arr := []rune(inp)

	for row := 0; row < gridLength; row++ {
		startIndex := row * gridLength
		rowSl[row] = string(arr[startIndex : startIndex+9])
		colSl := []rune(rowSl[row])
		for col := 0; col < gridLength; col++ {
			g[row][col] = int(colSl[col]) - 48
		}
	}

	return g

}

func prettyPrint(g [][]int) {

	gridLength := len(g)
	fmt.Printf("\n%v\n", strings.Repeat("-", 27))
	for row := 0; row < gridLength; row++ {
		fmt.Printf("| ")
		for col := 0; col < gridLength; col++ {
			fmt.Printf("%d ", g[row][col])
			if (col+1)%3 == 0 {
				fmt.Printf("| ")
			}
		}
		if (row+1)%3 == 0 {
			fmt.Printf("\n%v\n", strings.Repeat("-", 27))
		} else {
			fmt.Printf("\n")
		}
	}

}
