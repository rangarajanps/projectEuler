package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test(t *testing.T) {

	type testData struct {
		input    [][]int
		expected int
	}
	inp1 := [][]int{
		[]int{3, 0, 0, 0},
		[]int{7, 4, 0, 0},
		[]int{2, 4, 6, 0},
		[]int{8, 5, 9, 3},
	}

	inp2 := [][]int{
		[]int{75, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{95, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{17, 47, 82, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{18, 35, 87, 10, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{20, 04, 82, 47, 65, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{19, 01, 23, 75, 03, 34, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{88, 02, 77, 73, 07, 63, 67, 0, 0, 0, 0, 0, 0, 0, 0},
		[]int{99, 65, 04, 28, 06, 16, 70, 92, 0, 0, 0, 0, 0, 0, 0},
		[]int{41, 41, 26, 56, 83, 40, 80, 70, 33, 0, 0, 0, 0, 0, 0},
		[]int{41, 48, 72, 33, 47, 32, 37, 16, 94, 29, 0, 0, 0, 0, 0},
		[]int{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14, 0, 0, 0, 0},
		[]int{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57, 0, 0, 0},
		[]int{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48, 0, 0},
		[]int{63, 66, 04, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31, 0},
		[]int{04, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 04, 23},
	}

	tc1 := testData{input: inp1, expected: 23}
	tc2 := testData{input: inp2, expected: 1074}
	inp3 := create_input_from_file("matrix3.txt")
	tc3 := testData{input: inp3, expected: 7273}

	testCases := []testData{tc1, tc2, tc3}

	for i, test := range testCases {
		actual := FindMaximumPathSum(test.input)
		if actual != test.expected {
			t.Errorf("TestFindMaximumPathSum(%v) failed - \nWant: %d , whereas Got: %d \n", i, test.expected, actual)
		}
	}
}

func create_input_from_file(filename string) [][]int {

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return create_slice(txtlines)
}

func create_slice(inp []string) [][]int {

	g := make([][]int, len(inp))
	for i := range g {
		g[i] = make([]int, len(inp))
	}

	for i := 0; i < len(inp); i++ {
		line := strings.Split(inp[i], " ")
		for j := 0; j < len(line); j++ {
			val, err := strconv.Atoi(line[j])
			if err != nil {
				fmt.Println("Error in parsing the line ", err)
				return g
			}
			g[i][j] = val
		}
	}

	return g
}
