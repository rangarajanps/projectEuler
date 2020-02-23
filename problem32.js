function checkPanDigits(numAsStr) {
  for ( var i = 1; i <10; i++) {
    if ( numAsStr.indexOf(i) == -1 ) return false;
  }
  return true;
}

function pandigitalProducts() {

  var nList = [];
  var sum = 0;
  var str = "";
  var start = 0;
  var prod = 1;

  for ( var i = 2; i <= 60; i++ ) {

    if ( i < 10 ) {
      start = 1234;
    } else {
      start = 123;
    }
    
    for ( var j = start; j <= 10000/i; j++ ) {
      str = "".concat(i,j, i*j);
      if ( str.length < 9 ) continue;
      if ( checkPanDigits(str) ) {
        prod = i*j;
        if ( nList.indexOf(prod) == -1 ) {
          nList.push(prod);
          sum = sum + prod;
        }
      }
    }

  }
  
  return sum;
}

pandigitalProducts();
