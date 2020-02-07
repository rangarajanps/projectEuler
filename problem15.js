function latticePaths(gridSize) {
  
  var k = gridSize*2;
  var n = gridSize;
  var result =1 ;
		
		if ( n-k == 0 ) {
			k = n-k;
		}
		
		for ( var x = 0; x < k ; x++) {
			result = result * (n-x);
			result = result / (x+1);
		}
		
		return result;
}

latticePaths(2);
