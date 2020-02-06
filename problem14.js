function generateCollatzNumber(n) {
  if ( n % 2 == 0 ) {
    return n/2;
  } else {
    return 3*n + 1;
  }
}

function getLimit(number) {

  var counter = 1;
  while (number != 1 ) {
    number = generateCollatzNumber(number);
    counter++;
  }
  return counter;

}

function longestCollatzSequence(limit) {

  var maxVal = 2;
  var maxNum = 2;
  var temp = 0;

  for ( var i = 3; i <= limit; i++ ) {
    temp = getLimit(i);
    if ( temp > maxVal) {
      maxVal = temp;
      maxNum = i;
    }
  }

  return maxNum;
}

console.log(longestCollatzSequence(14));
