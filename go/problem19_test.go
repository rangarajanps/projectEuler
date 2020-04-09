package main

import (
	"testing"
)

func TestCountingSundays(t *testing.T) {

	type testData struct {
		fromYear     int
		toYear       int
		numOfSundays int
	}
	td1 := testData{
		fromYear:     1943,
		toYear:       1946,
		numOfSundays: 6,
	}
	td2 := testData{
		fromYear:     1995,
		toYear:       2000,
		numOfSundays: 10,
	}
	td3 := testData{
		fromYear:     1901,
		toYear:       2000,
		numOfSundays: 171,
	}

	testCases := []testData{td1, td2, td3}

	for _, tData := range testCases {
		actual := CountingSundays(tData.fromYear, tData.toYear)
		if actual != tData.numOfSundays {
			t.Errorf("TestCountingSundays failed for input (%d, %d) - Want: %d - whereas , Got: %d",
				tData.fromYear, tData.toYear, tData.numOfSundays, actual)
		}
	}

}
