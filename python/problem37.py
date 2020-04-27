import math


def truncatablePrimeSum(requiredNum):
    count = 0
    sum = 0
    num = 11
    while count < requiredNum:
        if isPrime(num) and isTruncatablePrime(num):
            count += 1
            sum += num
        num += 2

    return sum


def isTruncatablePrime(num):
    div = 10
    while div < num:
        if not isPrime(int(num // div)) or not isPrime(num % div):
            return False
        div *= 10

    return True


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
    print(truncatablePrimeSum(8))
