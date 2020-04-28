def largestPandigitalMultiples():
    maxPandigitNum = 1
    for i in range(1, 20000):
        temp = ""
        for j in range(1, 10):
            prod = i * j
            temp = temp + str(prod)
            if len(temp) < 9:
                continue
            elif len(temp) == 9:
                if isPandigital(temp) and int(temp) > maxPandigitNum:
                        maxPandigitNum = int(temp)
            else:
                break

    return maxPandigitNum


def isPandigital(numAsStr):
    digits = ["1", "2", "3", "4", "5", "6", "7", "8", "9"]
    for i in range(0, len(digits)):
        if numAsStr.find(digits[i]) == -1:
            return False

    return True


if __name__=="__main__":
    print(largestPandigitalMultiples())
