package main

import (
	"fmt"
	"testing"
)

func TestIsSolved(t *testing.T) {
	soln := `483921657967345821251876493548132976729564138136798245372689514814253769695417382`
	grid := genGrid(soln)
	prettyPrint(grid)
	if isSolved(grid) {
		fmt.Println("Congrats, you are doing great - you've solved the puzzle, keep rocking !!")
	} else {
		t.Errorf("Puzzle not solved")
	}
}
