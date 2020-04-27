def lowestDenominatorOfDigitCancelling():
    productNum = 1
    productDen = 1

    for a in range(10, 100):
        for b in range(a + 1, 100):
            if isDigitCancelling(a, b):
                productNum *= a
                productDen *= b

    greatestFactor = gcf(productNum, productDen)

    return productDen / greatestFactor


# Function that returns whether a fraction is a digit cancelling fraction a/b
def isDigitCancelling(a, b):
    # If they're both multiples of ten, return false
    if a % 10 == 0 and b % 10 == 0:
        return False

    # Get all the relevant digits isolated
    numTens = a // 10
    numOnes = a % 10
    denTens = b // 10
    denOnes = b % 10

    # Check all cases
    # numTens = denTens
    if numTens == denTens and isSameFraction(a, b, numOnes, denOnes):
        return True

    # numTens = denOnes
    if numTens == denOnes and isSameFraction(a, b, numOnes, denTens):
        return True

    # numOnes = denTens
    if numOnes == denTens and isSameFraction(a, b, numTens, denOnes):
        return True

    # numOnes = denOnes
    if numOnes == denOnes and isSameFraction(a, b, numTens, denTens):
        return True

    return False


# Function that checks if two fractions a/b and c/d are the same
def isSameFraction(a, b, c, d):
    return a * d == b * c


# Function that finds the GCF of two numbers
def gcf(a, b):
    smallest = a
    if a > b:
        smallest = b

    gcFactor = 1

    for i in range(1, smallest + 1):
        if a % i == 0 and b % i == 0:
            gcFactor = i

    return gcFactor


# Function to return the lowest denominator of fraction a/b
def lowestDenominator(a, b):
    greatestFactor = gcf(a, b)
    return b // greatestFactor


if __name__ == "__main__":
    print(lowestDenominatorOfDigitCancelling())
