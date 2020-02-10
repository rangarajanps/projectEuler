function maximumPathSumI(numTri) {
	var tmpTri = numTri;
	var botRow = tmpTri[tmpTri.length - 1];
	for(var i = 0; i < tmpTri.length-1; i++) {
		var checkRow = tmpTri[tmpTri.length - 2 - i];
		for(var j = 0; j < checkRow.length; j++) {
			if ( botRow[j] > botRow[j + 1] ) {
        checkRow[j] += botRow[j];
      } else {
        checkRow[j] += botRow[j + 1];
      }
		}
		botRow = checkRow;
	}

  return tmpTri[0][0];
}

const testTriangle = [[3, 0, 0, 0],
                      [7, 4, 0, 0],
                      [2, 4, 6, 0],
                      [8, 5, 9, 3]];

//maximumPathSumI(testTriangle);
