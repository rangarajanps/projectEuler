// Not an optimized solution, will not work beyond 5th power
// TODO - Understand how combinators work and refine this

function digitnPowers(n) {
  
  var init = parseInt("1".repeat(n));
  var limit = parseInt("9".repeat(n));
  if ( n >= 5 ) {
    init = 4000;
    limit = 199999;
  }
  //var limit = Math.pow(9,n);
  
  var temp;
  var sum = 0;
  var nList = [];
  for ( var i = init; i <= limit; i++ ) {
    
    var numToTest = i;
    temp = 0;
    if ( (""+numToTest).indexOf(5) != -1 || 
        (""+numToTest).indexOf(6) != -1 || 
        (""+numToTest).indexOf(7) != -1 || 
        (""+numToTest).indexOf(8) != -1 || 
        (""+numToTest).indexOf(9) != -1 ) {

      while ( numToTest > 0 ) {
        temp = temp + Math.pow(numToTest % 10, n);
        numToTest = parseInt(numToTest / 10);
      }
      if ( temp == i ) {
        nList.push(i);
        sum = sum + i;
      }

    }
    
  }
    
  return sum;
}

//console.log(digitnPowers(4));
