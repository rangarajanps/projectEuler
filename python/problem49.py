import math


def primePermutations():
    UNUSUAL_CONST = 3330

    p1 = 1488

    while True:

        p2 = p1 + UNUSUAL_CONST
        p3 = p2 + UNUSUAL_CONST

        if isPrime(p1) and isPrime(p2) and isPrime(p3):

            if checkIfThreeNumbersAreFormedOfSameDigit(p1, p2, p3):
                return str(p1) + str(p2) + str(p3)

        p1 += 1

    return ""


def checkIfThreeNumbersAreFormedOfSameDigit(n1, n2, n3):
    num1AsStr = convertToString(n1)
    num2AsStr = convertToString(n2)
    num3AsStr = convertToString(n3)
    if (num1AsStr == num2AsStr and num3AsStr == num2AsStr and num1AsStr == num3AsStr):
        return True

    return False


def convertToString(num):
    digits = []
    while num > 0:
        digits.append(num % 10)
        num = num // 10

    digits.sort()
    strings = [str(integer) for integer in digits]
    result = "".join(strings)

    return result


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


if __name__ == "__main__":
    print(primePermutations())
