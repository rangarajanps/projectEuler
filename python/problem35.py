import math


def circularPrimeCount(limit):
    count = 5
    # 2, 3, 5, 7, 11 is added as Circular Prime so that loop can be started
    # from 13 to skip all single digit numbers and rotated number being same as
    # initial input number. And also check only for odd numbers

    for n in range(13, limit, 2):

        if isPrime(n):

            number = n
            digits = numberOfDigits(number)
            powTen = math.pow(10, digits - 1)

            isCircularPrime = True
            for i in range(0, digits - 1):
                rem = number % 10
                quo = number // 10
                number = rem * powTen + quo
                if not isPrime(number):
                    isCircularPrime = False
                    break

            if isCircularPrime:
                #print("Found a Circular prime ",number)
                count += 1

    return count


# Function to return the count of digits of n
def numberOfDigits(n):
    cnt = 0
    while n > 0:
        cnt += 1
        n //= 10
    return cnt


def isPrime(num):
    if num <= 1:
        return False

    if num == 2:
        return True

    if num % 2 == 0:
        return False

    for i in range(3, int(math.sqrt(num))+1, 2):
        if num % i == 0:
            return False

    return True

if __name__=="__main__":
    print(circularPrimeCount(100))
