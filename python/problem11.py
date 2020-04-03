def get(arr, row, col):
    if row >= 0 and row < len(arr) and col >= 0 and col < len(arr):
        return arr[row][col]
    return 0

def max(numList):
    result = 1
    for val in numList:
        if val > result: result = val
    return result

def findLargestProduct(grid):

    maxProduct = 1

    for rowItr in range(0, len(grid)):
        for colItr in range(0, len(grid)):
            p1 = 1
            p2 = 1
            p3 = 1
            p4 = 1
            for i in range(0,4):
                p1 *= get(grid, rowItr, colItr+i)
                p2 *= get(grid, rowItr+i, colItr)
                p3 *= get(grid, rowItr+i, colItr+i)
                p4 *= get(grid, rowItr+i, colItr-i)
            product = max([p1,p2,p3,p4])
            if product > maxProduct: maxProduct=product
    return maxProduct

if __name__=='__main__':
    testGrid = [
        [40, 17, 81, 18, 57],
        [74, 4, 36, 16, 29],
        [36, 42, 69, 73, 45],
        [51, 54, 69, 16, 92],
        [7, 97, 57, 32, 16]
    ];
    print(findLargestProduct(testGrid))
