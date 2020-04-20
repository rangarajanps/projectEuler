import math


def digitNPower(n):
    init = int("1" * (n-1))
    limit = int("9" * (n+1))
    if init == 1:
        init +=1
    result = 0
    for i in range(init, limit+1):
        sum = sumOfDigitsPowerN(i, n)
        if sum == i:
            result += i
    return int(result)


def sumOfDigitsPowerN(num, n):
    sum = 0
    original_num = num
    while (num > 0):
        sum += math.pow(num % 10, n)
        if sum > original_num:
            return sum
        num = num // 10
    return sum


if __name__ == "__main__":
    print(digitNPower(2))
