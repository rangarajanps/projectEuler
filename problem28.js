function printArrayAsMatrix(arr) {
  var arrText='';
  for (var i = 0; i < arr.length; i++) {
    for (var j = 0; j < arr[i].length; j++) {
      arrText+=arr[i][j]+' ';
    }
    console.log(arrText);
    arrText='';
  }
}

function spiralDiagonals(limit) {
  
  var num = 1;
  var x = parseInt(limit/2);  
  var y = x;
  var grid = new Array(limit).fill(0).map(() => new Array(limit).fill(0)) 
  //console.log("Initial Grid is ");
  //printArrayAsMatrix(grid);

  grid[x][y] = num;
  num++;

  for ( var i = 0; i < limit/2; i++ ) {
    // Move Right i positions
    for ( var h = 0; h < 2*i+1; h++ ) {
      y = y + 1;    
      grid[x][y] = num;
      //console.log("Value in FOR Right at ["+x+" , "+y+"] is "+grid[x][y]);  
      num++;
    }
    if ( num >= limit*limit ) {
      break;
    }
    // Move Down i positions
    for ( var v = 0; v < 2*i+1; v++ ) {
      x = x + 1;
      grid[x][y] = num;
      //console.log("Value in FOR Down at ["+x+" , "+y+"] is "+grid[x][y]);
      num++;
    }
    // Move Left i positions
    for ( var h = 0; h < 2*i+2; h++ ) {
      y = y - 1;
      grid[x][y] = num;
      //console.log("Value in FOR Left at ["+x+" , "+y+"] is "+grid[x][y]);
      num++;
    }
    // Move Up i positions
    for ( var v = 0; v < 2*i+2; v++ ) {
      x = x - 1;
      grid[x][y] = num;
      //console.log("Value in FOR Up at ["+x+" , "+y+"] is "+grid[x][y]);
      num++;
    }
    //console.log("Grid AFTER iteration "+i+" is ");
    //printArrayAsMatrix(grid);
  }
  //console.log("Final Grid is ");
  //printArrayAsMatrix(grid);

  var sum = 0;
  
  for ( var p = 0; p < limit; p++ ) {
      sum = sum + grid[p][p];
  }
  //console.log("Sum of forward diagonal is "+sum);
  for ( var p = limit-1; p >= 0; p-- ) {
    if ( p != (limit-p-1) ) {
      sum = sum + grid[limit-p-1][p];
    }
  }

  return sum;
}

console.log(spiralDiagonals(5));
