def pandigitNumbers():
    prodSet = set()
    for i in range(2, 101):
        for j in range(100, 10001):
            prod = i * j
            pandigitNum = str(i) + str(j) + str(prod)
            if (len(pandigitNum) != 9): continue
            if checkPandigitNumber(pandigitNum):
                prodSet.add(prod)

    sum = 0
    for number in prodSet:
        sum += number

    return sum


def checkPandigitNumber(numAsStr):
    nums = list(numAsStr)
    for i in range(1, 10):
        if not ( str(i) in nums):
            return False
    return True


if __name__ == "__main__":
    print(pandigitNumbers())
