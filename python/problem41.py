import math

def largestNLengthPandigitalPrime(n):

	startIndex = int(math.pow(10, n-1)) + 1
	endIndex = int(math.pow(10,n)) - 1

	maxNLengthPandigitNum = 1
	for i in range(startIndex, endIndex):

		if isPandigital(i, n) and isPrime(i):
			if i > maxNLengthPandigitNum:
				maxNLengthPandigitNum = i

	return maxNLengthPandigitNum


def isPandigital(num, uptoN):
    numAsStr = str(num)
    digits = ["1", "2", "3", "4", "5", "6", "7", "8", "9"]
    for i in range(0, uptoN):
        if numAsStr.find(digits[i]) == -1:
            return False

    return True


def isPrime(num):
    if num <= 1:
        return False

    if num == 2:
        return True

    if num % 2 == 0:
        return False

    for i in range(3, int(math.sqrt(num)) + 1, 2):
        if num % i == 0:
            return False

    return True


if __name__=="__main__":
    print(largestNLengthPandigitalPrime(4))
