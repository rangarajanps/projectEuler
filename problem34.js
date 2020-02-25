// TODO - Should optimize further - have done brute force digits check
// it should be possible to do combinatorix digit selection and make it run from UI also

function generateFactorials() {

  var facts = [1,1,2];
  var fact = 1;
  for (var i = 3; i <= 9; i++) {
    for (var j = 2; j <= i; j++ ) {
      fact = fact * j;
    }
    facts[i] = fact;
    fact = 1;
  }
  return facts;
}

function isDigitFactorial(num) {
    
  var numToTest = num;
  var sum = 0;
  var rem = 0;

  var str = ""+num;
  if ( str.length == 3 && num >= 720 ) return false;
  if ( str.length == 3 ) {
    if ( str.indexOf(6) != -1 ) return false;
    if ( str.indexOf(7) != -1 ) return false;
    if ( str.indexOf(8) != -1 ) return false;
    if ( str.indexOf(9) != -1 ) return false;  
  }
  if ( str.length == 4 ) {
    if ( str.indexOf(8) != -1 ) return false;
    if ( str.indexOf(9) != -1 ) return false;  
  }
  if ( str.indexOf(9) != -1 ) return false;

  while ( numToTest >= 1 ) {
    rem = numToTest % 10;
    if ( sum > num || factorials[rem] >= num ) return false;
    sum = sum + factorials[rem];
    numToTest = parseInt(numToTest / 10);
  }
  if ( sum == num ) {
    return true;
  }
  return false;

}

function digitFactorial() {
  
  var sum = 0;
  var numbers = [];

  for ( var i = 100; i <= 40888; i++ ) {
    if ( isDigitFactorial(i) ) {
      //console.log("Found a Digit Factorials "+i);
      sum = sum + i;
      numbers.push(i)
    }
  }
  return { sum, numbers };
}

var factorials = generateFactorials();

digitFactorial();
