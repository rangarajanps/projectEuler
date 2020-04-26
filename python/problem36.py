def doubleBasePalindromeSum(limit):
    sum = 0
    for i in range(1, limit + 1):
        if isPalindrome(i) and isBinaryPalindrome(i):
            sum += i

    return sum


def isPalindrome(n):
    origNumber = n
    reversedNum = 0
    while n > 0:
        reversedNum = reversedNum * 10 + n % 10
        n = n // 10

    return reversedNum == origNumber


def isBinaryPalindrome(n):
    numInBase2 = []
    itr = 0
    while n > 0:
        numInBase2.append(n % 2)
        n = n // 2
        itr += 1

    for i in range(0, len(numInBase2) // 2):
        if numInBase2[i] != numInBase2[len(numInBase2) - i - 1]:
            return False

    return True


if __name__ == "__main__":
    print(doubleBasePalindromeSum(1000))
