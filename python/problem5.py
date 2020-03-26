def findSmallestFactor(limit):
    startIndex = generateSmallestFactor(limit)
    allDivisible = True
    while (allDivisible ) :
        for i in range(2,limit+1):
            if startIndex%i != 0:
                allDivisible = False
                break
        if not allDivisible:
            allDivisible=True
            startIndex += 1
        else:
            return startIndex


def generateSmallestFactor(limit):
    allFactorSet = set()
    for n in range(2, limit+1):
        fset = findFactors(n)
        allFactorSet.update(fset)
    print(allFactorSet)
    result = 1
    for factor in allFactorSet: result *= factor
    print(result)
    return result

def findFactors(num):
    factorSet = set()
    if num%2 == 0:
         factorSet.add(2)

    for x in range(2, num//2):
        if num%x == 0: factorSet.add(x)
    if len(factorSet) == 0: factorSet.add(num)
    return factorSet

if __name__ == '__main__':
    print(findSmallestFactor(13))
