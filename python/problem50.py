def consecutivePrimeSum(uptoN):
    isPrime = {}
    for i in range(2, uptoN + 1):
        isPrime[i] = True

    i = 2
    while (i * i <= uptoN):
        if (isPrime[i]):
            j = 2
            while (i * j <= uptoN):
                isPrime[i * j] = False
                j += 1
        i += 1

    primeNumbers = []
    for i in range(2, len(isPrime) + 1):
        if isPrime[i]:
            primeNumbers.append(i)

    finalSum = 0
    finalCount = 0

    for start in range(0, len(primeNumbers)):
        sum = 0
        count = 0

        for current in range(start, len(primeNumbers)):

            sum += primeNumbers[current]
            count += 1
            if sum >= uptoN:
                break

            if isPrime[sum] and count > finalCount:
                finalCount = count
                finalSum = sum

    return finalSum

if __name__ == "__main__":
    print(consecutivePrimeSum(1000))
