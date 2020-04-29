import math


def intRightTriangles(n):
    maxPeri = 0
    maxCount = 0
    temp = 0

    for i in range(8, n + 1):
        temp = countRightTrianglesWithPeri(i)
        if temp >= maxCount:
            maxCount = temp
            maxPeri = i

    return maxPeri


def countRightTrianglesWithPeri(n):
    count = 0
    c = 0
    for a in range(11, n // 3 + 1):
        for b in range(a + 1, n // 2 + 1):
            temp = math.sqrt(a * a + b * b)
            if not temp.is_integer():
                continue

            c = int(temp)
            if a + b + c == n:
                count += 1

    return count


if __name__ == "__main__":
    print(intRightTriangles(500))
