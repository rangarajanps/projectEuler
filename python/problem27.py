import math

def quadraticPrime(limit):
    maxA, maxB, maxN = 0, 0, 0

    for a in range(-limit, limit):

        for b in range(-limit, limit):

            count = 0
            while count < b:
                if isPrime(quadraticVal(a, b, count)):
                    count += 1
                else:
                    break

            if count > maxN:
                maxN = count
                maxA = a
                maxB = b

    return maxA * maxB


def quadraticVal(a, b, n):
    return int(abs(n * n + a * n + b))


def isPrime(number):
    if number <= 1:
        return False

    if number % 2 == 0:
        return False

    for i in range(3, int(math.sqrt(number)), 2):
        if number % i == 0:
            return False

    return True


if __name__ == "__main__":
    print(quadraticPrime(100))
