import math


def isPentagonal(n):
    p = (1 + math.sqrt(1 + 24 * n)) / 6
    return p == math.floor(p)


def isHexagonal(n):
    p = (1 + math.sqrt(1 + 8 * n)) / 4
    return p == math.floor(p)


def triPentaHexa():
    numToTest = 285
    result = 0
    while (True):

        numToTest += 1
        result = numToTest * (2 * numToTest - 1)

        if (isPentagonal(result) and isHexagonal(result)):
            return result


if __name__ == "__main__":
    print(triPentaHexa())
