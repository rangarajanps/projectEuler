daysInMonth = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
monthNames = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"]
dayInWeek = ["Su", "M", "Tu", "W", "Th", "F", "Sa"]

def isLeap(year):

	if year%4 == 0:
		if year%100 != 0:
			return True

		if year%400 == 0:
			return True

	return False

def findDayOf1stJanOfAYear(firstYear):

	day = 0 #1899 Jan 1st is Sunday
	for i in range(1900,firstYear+1):
		day = day + 1
		if isLeap(i):
			day = day + 1

	return (day % 7)

def countingSundays(firstYear, lastYear):

	day = 0
	countOfSundays = 0
	firstOfAYear = findDayOf1stJanOfAYear(firstYear)

	for i in range(firstYear,lastYear+1):

		day = day + firstOfAYear

		for j in range(0,len(daysInMonth)):

			if day%7 == 0:
				countOfSundays += 1

			day = day + daysInMonth[j]
			if j == 1 and isLeap(i):
				day = day + 1

		day = 0
		firstOfAYear = firstOfAYear + 1
		if isLeap(i):
			firstOfAYear = firstOfAYear + 1

	return countOfSundays

if __name__=='__main__':
    print(countingSundays(2000,2010))
