function findDivisor(n) {

  var divList = [];
  for ( var i = 1; i <= n/2; i++) {
    if ( n % i == 0 ) {
      divList.push(i);
    }
  }

  return divList;

}

function sumAmicableNum(n) {
  
  var divisors = [];
  var temp = 0;
  var properDiv = [];
  properDiv.push(0);

  for ( var p = 1; p <= n; p++ ) {

    divisors = findDivisor(p);
    temp = 0;
    for ( var i = 0; i < divisors.length; i++ ) {
        temp = temp + divisors[i];
    }
    properDiv.push(temp);
  }
  //console.log("Div sum list length is "+properDiv.length);
  //console.log("Div sum list is "+properDiv);
  
  var amicSum = 0;
  var amicList = [];
  for ( var x = 2; x < properDiv.length; x++ ) {
    
      for ( var y = 2; y < x; y++ ) {
        if ( x != y && properDiv[x] > 1 && x == properDiv[y] && y == properDiv[x] ) {
          //console.log("Found an amicable pair "+x+" "+y+" "+properDiv[x]+" "+properDiv[y]);
          amicList.push(x);
          amicList.push(y);
          amicSum = amicSum + x + y;
        }
      }
  }

  return amicSum;
}

sumAmicableNum(285);
