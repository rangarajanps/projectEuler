def sieve(limit):
    isPrimeList = []
    for i in range(1,limit+1):
        isPrimeList.append(True)

    for i in range(2, limit//2):
        j = 2
        while i*j < limit:
            isPrimeList[i*j]=False
            j+=1

    primes = []
    for i in range(2,limit):
        if isPrimeList[i]: primes.append(i)

    return primes

def sumOfPrimesTill(n):

    primeNumbers = sieve(n)
    #print(primeNumbers)
    sum = 0
    for prime in primeNumbers: sum += prime
    return sum

if __name__=='__main__':
    print(sumOfPrimesTill(10))
