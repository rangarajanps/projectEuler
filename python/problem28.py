def spiralDiagonalsSum(limit):
    grid = [[0 for x in range(limit)] for y in range(limit)]

    x = limit // 2
    y = x
    num = 1

    for i in range(0, limit // 2 + 1):

        # Move Right i positions
        for h in range(0, 2 * i + 1):
            grid[x][y] = num
            y = y + 1
            num += 1

        if num >= limit * limit:
            break

        # Move Down i positions
        for v in range(0, 2 * i + 1):
            grid[x][y] = num
            x = x + 1
            num += 1

        # Move Left i positions
        for h in range(0, 2 * i + 2):
            grid[x][y] = num
            y = y - 1
            num += 1

        # Move Up i positions
        for v in range(0, 2 * i + 2):
            grid[x][y] = num
            x = x - 1
            num += 1

    sum = 0

    for p in range(0, limit):
        sum = sum + grid[p][p]

    for p in range(limit - 1, -1, -1):
        if p != (limit - p - 1):
            sum = sum + grid[limit - p - 1][p]

    return sum

if __name__=="__main__":
    print(spiralDiagonalsSum(5))
