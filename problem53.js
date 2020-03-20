function buildPascal(n){
  
	var pascal = [];
	pascal[0] = [1];

	for(var row = 1; row <= n; row++){
		pascal[row] = [];

		for(var col = 0; col <= row; col++){
			if(col == 0){
				pascal[row][col] = pascal[row-1][col];	
			}else if(col == row){
				pascal[row][col] = pascal[row-1][col-1];
			}else{
				pascal[row][col] = pascal[row-1][col-1]+pascal[row-1][col];
			}
		}
	}
	return pascal;
}


function combinatoricSelections(limit) {

  var count = 0;
  var pascal = buildPascal(100);
  for(var i = 0; i < pascal.length; i++){
	  for(var j = 0; j < pascal[i].length; j++){
		  if(pascal[i][j] > limit){
			  count++;
		  }
	  }
  }

  return count;
}

//combinatoricSelections(1000000);
