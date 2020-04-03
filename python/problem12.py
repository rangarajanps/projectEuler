import math

def countDivisors(num):
    incr = 1
    if num % 2 != 0: incr = 2
    count = 2
    if num % 2 == 0: count = 4
    for i in range(3, int(math.sqrt(num))+1, incr):
        if num%i == 0: count += 2
    return count

def findHighlyDivisibleTriangularNumber(limit):
    triNum = 1
    incr = 2
    while True:
        triNum += incr
        if countDivisors(triNum) >= limit: return triNum
        incr += 1

if __name__=='__main__':
    print(findHighlyDivisibleTriangularNumber(5))
