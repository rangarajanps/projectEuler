import math

def isPrime(num):
    if num == 1: return False
    for x in range(2,int(math.sqrt(num))+1):
        if num % x == 0: return False

    return True

def generateNthPrimeNumber(limit):
    count = 1
    n = 3
    while True:
        if isPrime(n): count += 1
        if count == limit: break
        n += 1
    return n

if __name__=='__main__':
    print(generateNthPrimeNumber(6))
