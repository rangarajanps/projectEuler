def latticePaths(gridSize):

    result = 1
    n = gridSize * 2
    k = gridSize

    if n - k == 0:
        k = n - k

    for x in range(0,k):
        result = result * (n - x)
        result = result / (x + 1)

    return result

if __name__=='__main__':
    print(latticePaths(2))
