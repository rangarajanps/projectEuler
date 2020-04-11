import math

def findAmicableNumberSum(limit):
    factorSumMap = {}
    for i in range(2, limit + 1):
        factorSumMap[i] = factorSum(i)

    amicableNumberSum = 0
    for key in factorSumMap:
        val = factorSumMap.get(key)
        if val and key == factorSumMap.get(val) and key != val:
            amicableNumberSum = amicableNumberSum + val + key
            factorSumMap[val] = -1  # to avoid duplicate count or need to divide by 2

    return amicableNumberSum


def factorSum(number):
    factorSum = 1
    incrementor = 1
    if number % 2 != 0:
        incrementor = 2
    else:
        quotient = number / 2
        if quotient != 2:
            factorSum += quotient

        factorSum += 2

    for i in range(3, int(math.sqrt(number)), incrementor):
        if number % i == 0:
            quotient = number // i
            if quotient != i:
                factorSum += quotient

            factorSum += i

    return factorSum


if __name__ == '__main__':
    print(findAmicableNumberSum(1000))
