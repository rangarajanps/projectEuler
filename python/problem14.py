def genCollatzNumber(n):
    if n%2 == 0:
        return n/2
    else:
        return 3*n + 1

def countCollatzSequence(num):
    count=1
    while num != 1:
        num = genCollatzNumber(num)
        count += 1

    return count

def findLongestCollatzSequence(limit):
    maxCount = 1
    maxNum = 1
    for i in range(2,limit):
        result = countCollatzSequence(i)
        if result > maxCount:
            maxCount = result
            maxNum = i

    return maxNum

if __name__=='__main__':
    print(findLongestCollatzSequence(14))
