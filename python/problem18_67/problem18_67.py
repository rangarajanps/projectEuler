def maximumPathSum(triangle):
    for row in range(len(triangle)-2,-1,-1):
        for col in range(0,row+1):
            v1 = triangle[row][col] + triangle[row+1][col]
            v2 = triangle[row][col] + triangle[row+1][col+1]
            if v1 > v2:
                triangle[row][col] = v1
            else:
                triangle[row][col] = v2

    return triangle[0][0]

if __name__=='__main__':
    m = [[3, 0, 0, 0],
		[7, 4, 0, 0],
		[2, 4, 6, 0],
		[8, 5, 9, 3]]
    print(maximumPathSum(m))