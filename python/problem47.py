import math


def distinctPrimeFactors(targetFactorCount, targetConsecutiveNumCount):
    consecutiveNumCount = 0
    factorCount = 0
    startingNumber = 0
    numToCheck = 10
    foundOne = False

    while (True):

        if isPrime(numToCheck):
            numToCheck += 1
            foundOne = False
            consecutiveNumCount = 0
            continue

        factorCount = calculatePrimeFactors(numToCheck)
        if factorCount == targetFactorCount:
            if not foundOne:
                foundOne = True
                startingNumber = numToCheck
                consecutiveNumCount = 1
            else:
                consecutiveNumCount += 1
                if consecutiveNumCount == targetConsecutiveNumCount:
                    return startingNumber
        else:
            foundOne = False
            consecutiveNumCount = 0
        numToCheck += 1


def calculatePrimeFactors(num):
    factors = []
    factorsCountMap = {}

    for i in range(2, int(math.sqrt(num))+1):

        if num % i == 0:
            numToTest = num // i
            if isPrime(i):
                updateMap(factorsCountMap, i)
            if isPrime(numToTest):
                updateMap(factorsCountMap, numToTest)

            if numToTest == i:
                continue

            while numToTest > 0:
                if numToTest % i == 0:
                    if isPrime(i):
                        updateMap(factorsCountMap, i)
                    numToTest = numToTest // i
                else:
                    break

    for key in factorsCountMap:
        if factorsCountMap[key] == 1:
            factors.append(key)
        else:
            product = 1
            for p in range(0, factorsCountMap[key]):
                product *= key
            factors.append(product)

    return len(factors)


def updateMap(fMap, factor):
    if factor in fMap:
        fMap[factor] = fMap[factor] + 1
    else:
        fMap[factor] = 1


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
    print(distinctPrimeFactors(3, 3))
