package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

func TestNameScores(t *testing.T) {
	type testData struct {
		input    []string
		expected int
	}
	test1 := []string{`THIS`, `IS`, `ONLY`, `A`, `TEST`}
	test2 := []string{`I`, `REPEAT`, `THIS`, `IS`, `ONLY`, `A`, `TEST`}
	tc1 := testData{input: test1, expected: 791}
	tc2 := testData{input: test2, expected: 1468}
	inp3 := create_slice_from_file("p022_names.txt")
	tc3 := testData{input: inp3, expected: 871198282}

	testCases := []testData{tc1, tc2, tc3}

	for index, testCase := range testCases {
		actual := NameScores(testCase.input)
		if actual != testCase.expected {
			t.Errorf("TestNameScores failed for %d : Want: %d whereas , Got: %d \n", index, testCase.expected, actual)
		}
	}
}

func create_slice_from_file(filename string) []string {

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

	// why this file? Taken from https://programming.guide/go/read-file-line-by-line.html
	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var inpList []string

	for i := 0; i < len(txtlines); i++ {
		line := strings.Split(txtlines[i], ",")
		inpList = append(inpList, line...)
	}

	return inpList
}
