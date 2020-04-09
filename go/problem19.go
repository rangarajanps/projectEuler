package main

var daysInMonth = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
var monthNames = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
var dayInWeek = []string{"Su", "M", "Tu", "W", "Th", "F", "Sa"}

func isLeap(year int) bool {

	if year%4 == 0 {
		if year%100 != 0 {
			return true
		}
		if year%400 == 0 {
			return true
		}
	}
	return false
}

func findDayOf1stJanOfAYear(firstYear int) int {

	var day = 0 //1899 Jan 1st is Sunday
	for i := 1900; i <= firstYear; i++ {
		day = day + 1
		if isLeap(i) {
			day = day + 1
		}
	}

	return (day % 7)
}

func CountingSundays(firstYear, lastYear int) int {

	var day = 0
	var countOfSundays = 0
	var firstOfAYear = findDayOf1stJanOfAYear(firstYear)

	for i := firstYear; i <= lastYear; i++ {

		day = day + firstOfAYear

		for j := 0; j < len(daysInMonth); j++ {

			if day%7 == 0 {
				countOfSundays++
			}
			day = day + daysInMonth[j]
			if j == 1 && isLeap(i) {
				day = day + 1
			}
		}

		day = 0
		firstOfAYear = firstOfAYear + 1
		if isLeap(i) {
			firstOfAYear = firstOfAYear + 1
		}
	}

	return countOfSundays
}
