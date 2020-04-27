def generateFactorials(limit):
    factMap = [1,1,2]
    fact = 2
    for i in range(3, limit + 1):
        fact = fact * i
        factMap.append(fact)

    return factMap


factorials = generateFactorials(9)


def digitFactorialsSum():
    result = 0
    for n in range(100, 40880):
        if isDigitFactorial(n):
            result += n
    return result


def isDigitFactorial(num):
    numToTest = num
    sum = 0
    rem = 0

    while numToTest >= 1:
        rem = numToTest % 10
        if sum > num or factorials[rem] >= num:
            return False

        sum = sum + factorials[rem]
        numToTest = numToTest // 10

    if sum == num:
        return True

    return False


if __name__=="__main__":
    print(digitFactorialsSum())