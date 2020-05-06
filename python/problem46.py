import math


def goldbachOtherConjecture():
    isPrime = [True for i in range(6000)]

    for i in range(2, 6000 // 2):
        j = 2
        while (i * j < 6000):
            isPrime[i * j] = False
            j += 1

    compNumbers = []
    primeNumbers = []
    for i in range(2, 6000):
        if isPrime[i]:
            primeNumbers.append(i)
        else:
            compNumbers.append(i)

    for i in range(0, len(compNumbers)):
        if compNumbers[i] % 2 == 0: continue
        if not isGoldbach(compNumbers[i], primeNumbers):
            return compNumbers[i]

    return -1


def isGoldbach(num, primeList):
    
    for i in range(0, len(primeList)):
        if (primeList[i] > num): return False
        temp = num - primeList[i]
        if temp % 2 != 0:
            continue
        sqrtVal = math.sqrt(temp // 2)
        if sqrtVal.is_integer():
            return True

    return False

if __name__=="__main__":
    print(goldbachOtherConjecture())
