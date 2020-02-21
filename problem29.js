function distinctPowers(n) {
  
  var sqList = [];
  var temp = 0;
  for ( var i = 2; i <= n ; i++ ) {
    for ( var j = 2; j <= n; j++ ) {
      temp = Math.pow(i,j);
      if ( sqList.indexOf(temp) == -1 ) {
        sqList.push(temp);
      }
    }
    
  }

  return sqList.length;
}

console.log(distinctPowers(30));
