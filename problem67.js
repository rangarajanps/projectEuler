function maximumPathSumII(triangle) {
  
  for (var row = triangle.length - 2; row >= 0; row--) {

		for (var col = 0; col < row+1; col++) {

			var v1 = triangle[row][col] + triangle[row+1][col]
			var v2 = triangle[row][col] + triangle[row+1][col+1]
			if (v1 > v2) {
				triangle[row][col] = v1
			} else {
				triangle[row][col] = v2
			}

		}
	}
	
  console.log(triangle[0][0])
  // console.log(triangle)
	return triangle[0][0]
}

const testTriangle = [[3, 0, 0, 0],
                      [7, 4, 0, 0],
                      [2, 4, 6, 0],
                      [8, 5, 9, 3]];

// maximumPathSumII(testTriangle);
