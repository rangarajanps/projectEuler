def findMaxPalindrome(digits):
    
    max = int("9"+"0"*(digits*2-1))
    startIndex = int("9"+"0"*(digits-1))
    endIndex = int("1" + "0" * (digits))

    for x in range(startIndex, endIndex):
        for y in range(startIndex, endIndex):
            temp = x*y
            if isPalindrome(temp) and temp > max: max = temp

    return max

def isPalindrome(num):

    numToReverse = num
    reverse = 0

    while numToReverse > 0:
        reverse = reverse*10 + numToReverse%10
        numToReverse = numToReverse // 10

    if reverse == num: return True

    return False

if __name__ == '__main__':
    print(findMaxPalindrome(2))
