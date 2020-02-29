function checkForPandigit(num) {

  var nStr = String(num);
  if (nStr.length != 9 ) return false;
  if ( nStr.indexOf(0) != -1 ) return false;
  for ( var i = 1; i <=9; i++ ) {
    if ( nStr.indexOf(i) == -1 ) return false;
  }
  return true;
}

function pandigitalMultiples() {
  var panDigit = 0;
  var panDList = [];
  
  var temp ;
  var prod ;
  for ( var i = 1; i<20000; i++ ) {
    temp = "";
    for ( var j = 1; j<10; j++) {
      prod = i*j;
      temp = temp+String(prod);
      if ( temp.length < 9 ) continue;
      if ( temp.length == 9 ) {
        if ( checkForPandigit(temp)) {
          if ( temp > panDigit ) panDigit = temp;
          panDList.push(panDigit);
          //console.log("Found a Pandigit "+panDigit);
        }
      }
      if ( temp.length > 9 ) break;
    }
  }

  return parseInt(panDigit);
}

console.log(pandigitalMultiples());
