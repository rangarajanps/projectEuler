def champernowneConstant(n):
    magnitude = 1
    product = 1

    while magnitude <= n:
        temp = nthDigitOfFractional(magnitude)
        product *= temp
        magnitude *= 10

    return product


# Function to find the nth digit (going from right to left) of a number p
def nthDigitOfInteger(n, p):
    while n > 0:
        p //= 10
        n -= 1

    return p % 10


# Function to find the nth digit of the fractional concatenation of the positive integers
def nthDigitOfFractional(n):
    currPlace = 1
    currNumDigits = 1

    while n > 0:
        for i in range(currPlace, currPlace * 10):
            n -= currNumDigits
            if n <= 0:
                return nthDigitOfInteger(abs(n), i)

        currNumDigits += 1
        currPlace *= 10

    return 0

if __name__=="__main__":
    print(champernowneConstant(100))
