import math


def isPentagonal(n):
    if (1 + (24 * n + 1) ** 0.5) % 6 == 0:
        return True
    return False


def solvePentagonal(n):
    return (n * (3 * n - 1)) / 2


def pentagonNumbers():
    i = 1

    while (True):
        for j in range(1, i):
            a = solvePentagonal(i)
            b = solvePentagonal(j)
            if isPentagonal(a + b) and isPentagonal(a - b):
                return a-b
        i += 1


if __name__ == "__main__":
    print(pentagonNumbers())
