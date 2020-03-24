def sieve(n):
    primes = []
    primeNumbers = []

    for x in range(1,n+1):
        primes.append(True)

    for x in range(2,n//2):
        y=2
        while x*y<n:
            primes[x*y]=False
            y +=1

    for x in range(1, n):
        if primes[x]:
            primeNumbers.append(x)

    return primeNumbers

def isPrime(n):
    if n==1: return False
    for x in range(2, int(n ** (1 / 2))+1):
        if n % x == 0: return  False
    return True;

def largestPrimeFactor(n):
    primeList = sieve(int(n ** (1 / 2))+1)
    if isPrime(n): return n

    largestFactor = 1
    for x in primeList:
        if n/x < 0: break
        if n % x == 0 and x > largestFactor:
            largestFactor = x

    return largestFactor

if __name__ == '__main__':
    print(largestPrimeFactor(13195))
