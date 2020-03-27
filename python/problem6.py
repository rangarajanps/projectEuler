def findSumSquareDifference2(limit):
    sumSquare = ((limit*(limit+1))//2)**2
    squareSum = sum(list([x *x for x in range(1, limit+1)]))
    
    #sq_list = [x * x for x in range(1, limit + 1)]
    #squareSum = sum(item for item in sq_list)
    
    return sumSquare-squareSum

def findSumSquareDifference(limit):
    sumSquare = ((limit*(limit+1))//2)**2
    squareSum = 0
    for x in range(1,limit+1):
        squareSum += x*x
    return sumSquare-squareSum

if __name__ == '__main__':
    print(findSumSquareDifference(10))
    print(findSumSquareDifference2(10))
